package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Roadmap struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	Name  string   `json:"name"`
	Links []string `json:"links"`
}

var roadmapData = Roadmap{
	Domains: []Domain{
		{
			Name: "Programming Language",
			Links: []string{
				"https://github.com/salmer/CppDeveloperRoadmap",
				"https://github.com/Alikhll/golang-developer-roadmap",
				"https://github.com/s4kibs4mi/java-developer-roadmap",
				"https://github.com/aliyr/Nodejs-Developer-Roadmap",
				"https://github.com/thecodeholic/php-developer-roadmap",
				"https://github.com/anshulrgoyal/rust-web-developer-roadmap",
			},
		},
		{
			Name: "Web Development",
			Links: []string{
				"https://github.com/sulco/angular-developer-roadmap",
				"https://github.com/saifaustcse/angular-developer-roadmap",
				"https://github.com/MoienTajik/AspNetCore-Developer-Roadmap",
				"https://github.com/kamranahmedse/developer-roadmap",
				"https://github.com/sadanandpai/frontend-learning-kit/blob/main/public/2024_FE_roadmap.pdf",
				"https://github.com/Hasnayeen/laravel-developer-roadmap",
				"https://github.com/adam-golab/react-developer-roadmap",
				"https://github.com/flaviocopes/vue-developer-roadmap",
			},
		},
		{
			Name: "Mobile Development",
			Links: []string{
				"https://github.com/anacoimbrag/android-developer-roadmap",
				"https://github.com/olexale/flutter_roadmap",
				"https://github.com/BohdanOrlov/iOS-Developer-Roadmap",
			},
		},
		{
			Name: "Game Development",
			Links: []string{
				"https://github.com/utilForever/game-developer-roadmap",
				"https://github.com/miloyip/game-programmer",
			},
		},
		{
			Name: "AI / Machine Learning / Data Science",
			Links: []string{
				"https://github.com/AMAI-GmbH/AI-Expert-Roadmap",
				"https://github.com/floodsung/Deep-Learning-Papers-Reading-Roadmap",
				"https://github.com/machinelearningmindset/deep-learning-roadmap",
				"https://github.com/datastacktv/data-engineer-roadmap",
				"https://github.com/hasbrain/data-science-roadmap",
				"https://github.com/MrMimic/data-scientist-roadmap",
				"https://github.com/graykode/nlp-roadmap",
			},
		},
		{
			Name: "College Coursework",
			Links: []string{
				"https://drive.google.com/drive/folders/1TTjSSShjK5nrUURnSIVOK1wnO-CBCp7R",
				"https://drive.google.com/drive/folders/1c2gZ8TJ53P3ecNETMB403lTueQY5pEHU",
				"https://drive.google.com/drive/folders/1jUHgBPZdN5V5Am6gEHpHNTc4XWIQgC_W",
				"https://your-fourth-year-link-here.com", // Replace with the actual link
			},
		},
		{
			Name: "Miscellaneous",
			Links: []string{
				"https://github.com/fityanos/awesome-quality-assurance-roadmap",
				"https://github.com/Sundowndev/hacker-roadmap",
				"https://github.com/AlaaAttya/software-architect-roadmap",
				"https://github.com/stemmlerjs/software-design-and-architecture-roadmap",
				"https://github.com/mohsenshafiei/system-design-master-plan",
				"https://github.com/togiberlin/ui-ux-designer-roadmap",
				"https://github.com/nietsymerej/collecobrary",
			},
		},
	},
}

func handleRoadmap(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")

	if domain == "" {
		http.Error(w, "Missing domain parameter", http.StatusBadRequest)
		return
	}

	for _, d := range roadmapData.Domains {
		if d.Name == domain {
			jsonData, err := json.Marshal(d.Links)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
			return
		}
	}

	http.Error(w, "Domain not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/roadmap", handleRoadmap)
	fmt.Println("Port starting at :8080")
	http.ListenAndServe(":8080", nil)
}
