package internal

import (
    "net/http"
    "html/template"
)

// Viewer represents the document viewing area with zoom controls.
type Viewer struct {
    ZoomLevel int // Current zoom level
}

// NewViewer initializes a new Viewer with default zoom level.
func NewViewer() *Viewer {
    return &Viewer{ZoomLevel: 100}
}

// ServeHTTP serves the HTML page with zoom controls.
func (v *Viewer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.New("viewer").Parse(html))
    tmpl.Execute(w, v)
}

var html = `<!DOCTYPE html>
<html>
<head>
    <title>Document Viewer</title>
    <script>
        function zoomIn() {
            // Logic to increase zoom level
        }
        function zoomOut() {
            // Logic to decrease zoom level
        }
    </script>
</head>
<body>
    <h1>Document Viewer</h1>
    <div>
        <button onclick="zoomOut()">Zoom Out</button>
        <button onclick="zoomIn()">Zoom In</button>
    </div>
    <div>
        <p>Current Zoom Level: {{.ZoomLevel}}%</p>
    </div>
</body>
</html>`
