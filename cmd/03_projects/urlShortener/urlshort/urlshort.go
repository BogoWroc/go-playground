package urlshort

import (
	yamlV2 "gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	f := func(rw http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(rw, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(rw, r)
	}
	return f
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yaml []byte) (dst []map[string]string, err error) {
	err = yamlV2.Unmarshal(yaml, &dst)
	return dst, err
}

func buildMap(parsedYaml []map[string]string) map[string]string {
	mergedMap := make(map[string]string)
	for _, entry := range parsedYaml {
		key := entry["path"]
		mergedMap[key] = entry["url"]
	}
	return mergedMap
}
