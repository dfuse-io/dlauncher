// Copyright 2019 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dashboard

//go:generate rice embed-go
import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

func (s *server) dashboardStaticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}

	if path == "/" || path == "/index.html" {
		zlog.Debug("serving templated 'index.html'")
		s.serveIndexHTML(w, r)
		return
	}

	zlog.Debug("serving dashboard static asset", zap.String("path", path))
	pathFile, err := s.box.Open(path)
	if err != nil {
		zlog.Debug("static asset not found, falling back to 'index.html'")
		s.serveIndexHTML(w, r)
		return
	}
	defer pathFile.Close()

	stat, err := pathFile.Stat()
	if err != nil {
		zlog.Warn("cannot stat file, serving without a MIME type", zap.String("path", path))
		io.Copy(w, pathFile)
		return
	}

	http.ServeContent(w, r, path, stat.ModTime(), pathFile)
}

func (s *server) serveIndexHTML(w http.ResponseWriter, r *http.Request) {
	zlog.Debug("serving templated index.html'")

	w.Header().Set("Content-Type", "text/html")
	w.Write(s.indexData)
}
