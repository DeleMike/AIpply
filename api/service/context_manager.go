package service

import (
	"regexp"
	"strings"
)

type JobContext struct {
	Role     string
	Company  string
	Keywords []string
}

func ExtractJobContext(jobDescription string) JobContext {
	var ctx JobContext

	// Extract role (first capitalized phrase ending with "Engineer", "Manager", etc.)
	roleRegex := regexp.MustCompile(`(?i)([A-Z][a-zA-Z ]+(Engineer|Developer|Manager|Analyst|Designer|Consultant|Scientist))`)
	match := roleRegex.FindStringSubmatch(jobDescription)
	if len(match) > 0 {
		ctx.Role = strings.TrimSpace(match[0])
	}

	// Extract company (look for "at <Company>")
	companyRegex := regexp.MustCompile(`(?i)at\s+([A-Z][A-Za-z0-9&\s]+)`)
	match = companyRegex.FindStringSubmatch(jobDescription)
	if len(match) > 1 {
		ctx.Company = strings.TrimSpace(match[1])
	}

	// Extract domain keywords (basic set: Android, ML, Finance, Cloud, etc.)
	keywords := []string{}
	for _, kw := range []string{"Android", "iOS", "ML", "AI", "Finance", "Backend", "Frontend", "Cloud", "DevOps", "Security"} {
		if strings.Contains(strings.ToLower(jobDescription), strings.ToLower(kw)) {
			keywords = append(keywords, kw)
		}
	}
	ctx.Keywords = keywords

	return ctx
}
