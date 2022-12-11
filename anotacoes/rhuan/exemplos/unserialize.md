```golang
payload := rdts.NewTable()
err := json.NewDecoder(request.Body).Decode(&payload)
if err != nil {
	log.Println("error on json decode:", err)
	apiReturn(http.StatusUnprocessableEntity)
	return
}
```
