package aarm

import "net/http"

// SetHandles pode definir todos os handles de uma única vez.
func (api *Api) SetHandles(handles map[string]func(http.ResponseWriter, *http.Request)) {

	for endpoint, hundle := range handles {
		http.HandleFunc(endpoint, hundle)
	}

}
