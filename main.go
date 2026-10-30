package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

type Project struct {
	Name        string
	Description string
	URL         string
	Tags        []string
}

type Portfolio struct {
	Name         string
	Role         string
	Introduction string
	Projects     []Project
}

var portfolio = Portfolio{
	Name:         "Zidni",
	Role:         "Software engineer",
	Introduction: "I build reliable, thoughtful software with Go and React. Here are a few projects I have been working on.",
	Projects: []Project{
		{Name: "DataSync Engine", Description: "A resilient data synchronization pipeline with an outbox, snapshots, and stream processing.", URL: "https://github.com/zidni0192/datasync-engine", Tags: []string{"Go", "Redis Streams", "ClickHouse"}},
		{Name: "Multi-tenant SaaS", Description: "A tenant-aware application focused on isolation, role-based access, and auditability.", URL: "https://github.com/zidni0192/multitenant-saas", Tags: []string{"Go", "React", "PostgreSQL"}},
		{Name: "Go Job Scheduler", Description: "A durable job scheduler with retries, leases, and observable execution.", URL: "https://github.com/zidni0192/go-job-scheduler", Tags: []string{"Go", "SQLite", "Testing"}},
	},
}

var page = template.Must(template.New("page").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="description" content="{{.Introduction}}"><title>{{.Name}} — {{.Role}}</title>
  <style>
    :root{color-scheme:light;--ink:#16211f;--muted:#52615d;--paper:#f5f4ee;--card:#fffefa;--line:#d9ded5;--accent:#176b52;--soft:#e3eee6}
    *{box-sizing:border-box}body{margin:0;background:var(--paper);color:var(--ink);font:16px/1.6 system-ui,-apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif}
    main{width:min(100% - 2rem,880px);margin:0 auto;padding:5rem 0 4rem}.eyebrow{color:var(--accent);font-size:.8rem;font-weight:750;letter-spacing:.14em;text-transform:uppercase}
    h1{max-width:700px;margin:.5rem 0 1rem;font-size:clamp(2.5rem,8vw,5rem);line-height:1.04;letter-spacing:-.055em}h2{margin:0;font-size:1.35rem;letter-spacing:-.02em}
    .intro{max-width:650px;color:var(--muted);font-size:1.15rem}.section-head{display:flex;align-items:baseline;justify-content:space-between;margin:4rem 0 1rem}.section-head p{color:var(--muted);font-size:.9rem}
    .projects{display:grid;grid-template-columns:repeat(auto-fit,minmax(min(100%,250px),1fr));gap:1rem}.card{display:flex;flex-direction:column;min-height:230px;padding:1.35rem;border:1px solid var(--line);border-radius:14px;background:var(--card)}
    .card p{margin:.6rem 0 1rem;color:var(--muted)}.tags{display:flex;flex-wrap:wrap;gap:.45rem;margin-top:auto;padding:0;list-style:none}.tags li{padding:.2rem .55rem;border-radius:999px;background:var(--soft);color:#24563f;font-size:.78rem;font-weight:650}
    a{color:var(--accent);font-weight:700;text-decoration-thickness:1px;text-underline-offset:3px}a:hover{text-decoration-thickness:2px}a:focus-visible{outline:3px solid #d48334;outline-offset:4px;border-radius:2px}
    footer{margin-top:3rem;padding-top:1.2rem;border-top:1px solid var(--line);color:var(--muted);font-size:.9rem}@media(max-width:520px){main{padding-top:3.25rem}.section-head{display:block;margin-top:3rem}}
  </style>
</head>
<body><main>
  <header><p class="eyebrow">{{.Role}}</p><h1>Hello, I’m {{.Name}}.</h1><p class="intro">{{.Introduction}}</p></header>
  <section aria-labelledby="projects-heading"><div class="section-head"><h2 id="projects-heading">Selected projects</h2><p>{{len .Projects}} projects</p></div>
    {{if .Projects}}<div class="projects">{{range .Projects}}<article class="card"><h2>{{.Name}}</h2><p>{{.Description}}</p><ul class="tags" aria-label="Technologies">{{range .Tags}}<li>{{.}}</li>{{end}}</ul><p><a href="{{.URL}}" target="_blank" rel="noopener noreferrer" aria-label="View {{.Name}} on GitHub (opens in a new tab)">View project <span aria-hidden="true">↗</span></a></p></article>{{end}}</div>{{else}}<p>No projects to show yet. Check back soon.</p>{{end}}
  </section><footer><p>Built with Go. <a href="https://github.com/zidni0192">Find me on GitHub</a>.</p></footer>
</main></body></html>`))

func newHandler(content Portfolio) http.Handler {
	projects := append([]Project(nil), content.Projects...)
	sort.SliceStable(projects, func(i, j int) bool { return projects[i].Name < projects[j].Name })
	content.Projects = projects
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := page.Execute(w, content); err != nil {
			log.Printf("render portfolio page: %v", err)
		}
	})
}

func main() {
	log.Println("portfolio-site listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", newHandler(portfolio)))
}
