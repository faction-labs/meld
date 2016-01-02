package rust

func (r *RustServer) getResponse(requestId int) (string, error) {
	// wait for response
	resp := ""
	for respId := range r.respChan {
		// TODO: add timeout
		if respId == requestId {
			resp = r.requests[respId]
			delete(r.requests, respId)
			break
		}
	}

	return resp, nil
}
