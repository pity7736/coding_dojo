package exercises

type Service struct {
	Dur  int
	Deps []string
}

type serviceResult struct {
	duration int
	path     []string
}

func FindCriticalPath(services map[string]Service) (int, []string) {
	memo := map[string]serviceResult{}

	var compute func(name string) serviceResult
	compute = func(name string) serviceResult {
		if result, ok := memo[name]; ok {
			return result
		}

		service := services[name]
		bottleneck := serviceResult{}

		for _, dep := range service.Deps {
			depResult := compute(dep)
			if depResult.duration > bottleneck.duration {
				bottleneck = depResult
			}
		}

		result := serviceResult{
			duration: bottleneck.duration + service.Dur,
			path:     append(bottleneck.path, name),
		}
		memo[name] = result
		return result
	}

	final := serviceResult{}
	for name := range services {
		result := compute(name)
		if result.duration > final.duration {
			final = result
		}
	}

	return final.duration, final.path
}
