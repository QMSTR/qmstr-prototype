package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	model "qmstr-prototype/qmstr/qmstr-model"
)

var closeServer chan interface{}

func handleQuitRequest(w http.ResponseWriter, r *http.Request) {
	// nothing to do except quit:
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Bye now.\n")
	Info.Printf("handleQuitRequest: quit request received.")
	closeServer <- nil
}

func handleSourceRequest(w http.ResponseWriter, r *http.Request) {
	Info.Printf("handleSourceRequest: processing a %s request", r.Method)
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		s, err := Model.GetSourceEntity(id)
		if err != nil {
			// no such entity, this is not a master server error, return an empty source entity
			Info.Printf("handleSourceRequest: %s - no such entity, returning an empty one", r.Method)
			s = model.SourceEntity{Path: "", Hash: ""}
		}
		b, err := json.Marshal(s)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		Info.Printf("handleSourceRequest: %s - response: %v", r.Method, string(b[:]))
		w.Write(b)
	case "POST":
		id := r.URL.Query().Get("id")
		decoder := json.NewDecoder(r.Body)
		var s model.SourceEntity
		err := decoder.Decode(&s)
		if err != nil {
			Info.Printf("handleSourceRequest: %s - error parsing request body", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if s.ID() != id {
			//strange:
			s.Path = id
		}
		err = Model.AddSourceEntity(s)
		if err != nil {
			Info.Printf("handleSourceRequest: %s - error adding source entity: %s", r.Method, err.Error())
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		Info.Printf("handleSourceRequest: %s - done", r.Method)
		// else: done
	case "DELETE":
		id := r.URL.Query().Get("id")
		s, err := Model.GetSourceEntity(id)
		if err != nil {
			// no such entity, this is not a master server error, return an empty source entity
			Info.Printf("handleSourceRequest: %s - no such entity, cannot delete it", r.Method)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		if err := Model.DeleteSourceEntity(s); err != nil {
			Info.Printf("handleSourceRequest: %s - error deleting entity", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// else: done
	case "PUT":
		id := r.URL.Query().Get("id")

		if _, err := Model.GetSourceEntity(id); err != nil {
			Info.Printf("handleSourceRequest: %s - no such entity, cannot modify it", r.Method)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var s model.SourceEntity
		err := decoder.Decode(&s)
		if err != nil {
			Info.Printf("handleSourceRequest: %s - error parsing request body", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if s.ID() != id {
			//strange:
			s.Path = id
		}

		if err := Model.ModifySourceEntity(s); err != nil {
			Info.Printf("handleSourceRequest: %s - error modifying source entity", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// else: done
	default:
		Log.Printf("handleSourceRequest: don't know how to handle a %s request", r.Method)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}
}

func handleDependencyRequest(w http.ResponseWriter, r *http.Request) {
	Info.Printf("handleDependencyRequest: processing a %s request", r.Method)
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		s, err := Model.GetDependencyEntity(id)
		if err != nil {
			// no such entity, this is not a master server error, return an empty source entity
			Info.Printf("handleDependencyRequest: %s - no such entity, returning an empty one", r.Method)
			s = model.DependencyEntity{Name: "", Hash: ""}
		}
		b, err := json.Marshal(s)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		Info.Printf("handleDependencyRequest: %s - response: %v", r.Method, string(b[:]))
		w.Write(b)
	case "POST":
		id := r.URL.Query().Get("id")
		decoder := json.NewDecoder(r.Body)
		var d model.DependencyEntity
		err := decoder.Decode(&d)
		if err != nil {
			Info.Printf("handleDependencyRequest: %s - error parsing request body", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if d.ID() != id {
			//strange:
			d.Name = id
		}
		err = Model.AddDependencyEntity(d)
		if err != nil {
			Info.Printf("handleDependencyRequest: %s - error adding dependency entity: %s", r.Method, err.Error())
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		Info.Printf("handleDependencyRequest: %s - done", r.Method)
		// else: done
	case "DELETE":
		id := r.URL.Query().Get("id")
		d, err := Model.GetDependencyEntity(id)
		if err != nil {
			// no such entity, this is not a master server error, return an empty source entity
			Info.Printf("handleDependencyRequest: %s - no such entity, cannot delete it", r.Method)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		if err := Model.DeleteDependencyEntity(d); err != nil {
			Info.Printf("handleDependencyRequest: %s - error deleting entity", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// else: done
	case "PUT":
		id := r.URL.Query().Get("id")

		if _, err := Model.GetDependencyEntity(id); err != nil {
			Info.Printf("handleDependencyRequest: %s - no such entity, cannot modify it", r.Method)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var d model.DependencyEntity
		err := decoder.Decode(&d)
		if err != nil {
			Info.Printf("handleDependencyRequest: %s - error parsing request body", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if d.ID() != id {
			//strange:
			d.Name = id
		}

		if err := Model.ModifyDependencyEntity(d); err != nil {
			Info.Printf("handleSourceRequest: %s - error modifying source entity", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	// else: done
	default:
		Log.Printf("handleDependencyRequest: don't know how to handle a %s request", r.Method)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}
}

func handleTargetRequest(w http.ResponseWriter, r *http.Request) {
	Info.Printf("handleTargetRequest: processing a %s request", r.Method)
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		t, err := Model.GetTargetEntity(id)
		if err != nil {
			// no such entity, this is not a master server error, return an empty source entity
			Info.Printf("handleTargetRequest: %s - no such entity, returning an empty one", r.Method)
			t = model.TargetEntity{Name: "", Hash: ""}
		}
		b, err := json.Marshal(t)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		Info.Printf("handleTargetRequest: %s - response: %v", r.Method, string(b[:]))
		w.Write(b)
	case "POST":
		id := r.URL.Query().Get("id")
		decoder := json.NewDecoder(r.Body)
		var t model.TargetEntity
		err := decoder.Decode(&t)
		if err != nil {
			Info.Printf("handleTargetRequest: %s - error parsing request body", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if t.ID() != id {
			//strange:
			t.Name = id
		}
		err = Model.AddTargetEntity(t)
		if err != nil {
			Info.Printf("handleTargetRequest: %s - error adding entity: %s", r.Method, err.Error())
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		Info.Printf("handleTargetRequest: %s - done", r.Method)
		// else: done
	case "DELETE":
		id := r.URL.Query().Get("id")
		e, err := Model.GetTargetEntity(id)
		if err != nil {
			// no such entity, this is not a master server error, return an empty source entity
			Info.Printf("handleTargetRequest: %s - no such entity, cannot delete it", r.Method)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		if err := Model.DeleteTargetEntity(e); err != nil {
			Info.Printf("handleTargetRequest: %s - error deleting entity", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// else: done
	case "PUT":
		id := r.URL.Query().Get("id")

		if _, err := Model.GetTargetEntity(id); err != nil {
			Info.Printf("handleTargetRequest: %s - no such entity, cannot modify it", r.Method)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var e model.TargetEntity
		err := decoder.Decode(&e)
		if err != nil {
			Info.Printf("handleTargetRequest: %s - error parsing request body", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if e.ID() != id {
			//strange:
			e.Name = id
		}

		if err := Model.ModifyTargetEntity(e); err != nil {
			Info.Printf("handleSourceRequest: %s - error modifying source entity", r.Method)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// else: done
	default:
		Log.Printf("handleTargetRequest: don't know how to handle a %s request", r.Method)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}
}

func handleReportRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	Info.Printf("handleReportRequest: creating report...")
	id := r.URL.Query().Get("id")
	t, err := Model.GetTargetEntity(id)
	if err != nil {
		// no such entity, this is not a master server error, return an empty source entity
		Info.Printf("handleReportRequest: %s - no such entity, returning an empty one", r.Method)
		t = model.TargetEntity{Name: "", Hash: ""}
	}
	report := CreateReport(t)
	result := fmt.Sprintf("{ \"report\": %s }", report)
	w.Write([]byte(result))
}

func handleHealthRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	Info.Printf("handleHealthRequest: reporting on heath status...")
	// For now no real check is done; Just tell that we are running.
	w.Write([]byte("{ \"running\": \"ok\" }"))
}

func handleLinkedTargetsRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	Info.Printf("handleLinkedTargetsRequest: return linked targets...")

	b, err := json.Marshal(Model.GetAllLinkedTargets())
	if err == nil {
		result := fmt.Sprintf("{ \"linkedtargets\" : %s}", string(b))
		w.Write([]byte(result))
	} else {
		Info.Printf("Error: %v", err)
		w.Write([]byte("{}"))
	}
}

func handleDumpRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	Info.Printf("handleDumpRequest: dump data model...")
	dumpModel := "{ \"sources\": %s, \"targets\": %s, \"dependencies\": %s }"
	srcs := ""
	targets := ""
	deps := ""

	b, err := json.Marshal(Model.GetAllSourceEntities())
	if err == nil {
		srcs = string(b)
	} else {
		Info.Printf("Error: %v", err)
	}

	b, err = json.Marshal(Model.GetAllTargetEntities())
	if err == nil {
		targets = string(b)
	} else {
		Info.Printf("Error: %v", err)
	}

	b, err = json.Marshal(Model.GetAllTargetEntities())
	if err == nil {
		deps = string(b)
	} else {
		Info.Printf("Error: %v", err)
	}

	w.Write([]byte(fmt.Sprintf(dumpModel, srcs, targets, deps)))
}

func handleReuseRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{ \"reuse compliant\": \"ok\" }"))
}

func startHTTPServer() chan string {
	address := ":9000"
	server := &http.Server{Addr: address}
	http.HandleFunc("/quit", handleQuitRequest)
	http.HandleFunc("/sources", handleSourceRequest)
	http.HandleFunc("/dependencies", handleDependencyRequest)
	http.HandleFunc("/targets", handleTargetRequest)
	http.HandleFunc("/report", handleReportRequest)
	http.HandleFunc("/health", handleHealthRequest)
	http.HandleFunc("/dump", handleDumpRequest)
	http.HandleFunc("/linkedtargets", handleLinkedTargetsRequest)
	http.HandleFunc("/reuse", handleReuseRequest)

	Info.Printf("starting HTTP server on address %s", address)
	channel := make(chan string)
	go func() {
		err := server.ListenAndServe()
		server = nil
		if err == http.ErrServerClosed {
			channel <- fmt.Sprintf("startHTTPServer: server closed.")
		} else if err != nil {
			channel <- fmt.Sprintf("startHTTPServer: exiting with error: %s", err.Error())
		} else {
			channel <- "startHTTPServer: retreating coordinatedly."
		}
	}()

	closeServer = make(chan interface{})
	go func() {
		<-closeServer
		Info.Printf("shutting down HTTP server on address %s", address)
		if server != nil {
			if err := server.Shutdown(context.Background()); err != nil {
				panic(err) // failure/timeout shutting down the server gracefully
			}
		} else {
			Log.Printf("stopHTTPServer: server shutdown requested, but server is not running.")
		}
		close(closeServer)
		closeServer = nil
	}()
	return channel
}
