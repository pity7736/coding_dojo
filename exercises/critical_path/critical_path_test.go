package criticalpath_test

import (
	"coding_dojo/exercises/critical_path"
	"reflect"
	"testing"
)

func TestCriticalPathFinderShould(t *testing.T) {
	t.Run("find the critical path", func(t *testing.T) {
		services := map[string]criticalpath.Service{
			"db":      {Dur: 5, Deps: []string{}},
			"cache":   {Dur: 2, Deps: []string{}},
			"api":     {Dur: 3, Deps: []string{"db", "cache"}},
			"gateway": {Dur: 4, Deps: []string{"api"}},
		}
		duration, criticalPath := criticalpath.FindCriticalPath(services)

		expectedDuration := 12
		expectedPath := []string{"db", "api", "gateway"}

		if duration != expectedDuration {
			t.Errorf("duration expected: %d, got %d", expectedDuration, duration)
		}

		if !reflect.DeepEqual(expectedPath, criticalPath) {
			t.Errorf("critical path expected: %v, got %v", expectedPath, criticalPath)
		}

	})
}
