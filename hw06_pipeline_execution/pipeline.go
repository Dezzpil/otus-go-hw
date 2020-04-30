package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(done In, in In) (out Out)

func CreateStage(f func(v I) I) Stage {
	return func(done In, in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
			for {
				select {
				case <-done:
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					out <- f(v)
				}
			}
		}()
		return out
	}
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	chs := make([]In, 0, len(stages))
	chs = append(chs, in)

	for i := 0; i < len(stages); i++ {
		stage, in := stages[i], chs[i]
		out := stage(done, in)
		chs = append(chs, out)
	}

	return chs[len(chs)-1]
}
