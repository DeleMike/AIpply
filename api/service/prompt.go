package service

// QuestionPrompt is to be passed to LLM
const QuestionPrompt = `
You are an expert hiring manager and interview coach with experience across a vast number of industries, from technology and finance to creative arts and skilled trades. Your task is to generate a set of insightful interview questions based *only* on the provided job description and the candidate's experience level.

Analyze the job description to identify the core responsibilities, required skills (both hard and soft), and the likely challenges of the role.

Candidate Experience Level: %s

Job Description:
---
%s
---

Guidelines:
- Generate questions that are open-ended and tailored specifically to the details in the job description.
- Create a mix of behavioral questions ("Tell me about a time...") and situational questions ("How would you handle a situation where...").
- Ensure the number of questions matches the candidate's experience level: 8 for New Grad, 7 for Mid-Level, and 6 for Senior.
- **IMPORTANT: Your entire response must consist ONLY of the questions.** Do not include any introductory text, commentary, numbering, or bullet points. Each question must be on a new line.
`

// CVPrompt is to be passed to LLM
const CVPrompt = `
You are an elite career coach and professional resume writer. Your task is to generate a complete, modern, and highly-tailored CV.

**You will receive two inputs:**
1.  **The Target Job Description:** The specific role the user is applying for.
2.  **The User's Answers:** A JSON array of question-and-answer pairs.

Your goal is to synthesize all this information into a single, cohesive CV document that is optimized to get the user an interview for the target job.

**Your Process:**

1.  **Analyze the Job Description:** First, read the job description to identify all key skills, technologies, and required qualifications (e.g., "React," "5+ years of experience," "team leadership," "data analysis").
2.  **Parse the User's Answers:** The answers are in a "[{"question": "...", "answer": "..."}]" format. Use the question to understand the context of each answer.
3.  **Extract Data:**
    * Find the user's name, email, and other contact details (from questions like "What's your full name and email?").
    * Find their work history (from questions like "What was your most recent job title and company?").
4.  **Transform Achievements (The Most Important Step):**
    * Find all the behavioral/STAR-method answers (from questions like "Tell me about a time..." or "Describe a project...").
    * **Convert these narrative stories into concise, powerful, action-oriented bullet points** for the "Work Experience" section.
    * **Quantify results** whenever possible (e.g., if the user says "I made the app faster," transform it into "Increased app performance by 30%%...").
5.  **Build the CV Sections:**
    * **Contact Details:** (Name, Email, Phone, LinkedIn placeholder).
    * **Professional Summary:** Write a 2-3 sentence summary at the top that perfectly aligns the user's strongest skills (from their answers) with the target job description.
    * **Skills:** Create a skills list using a mix of skills from the user's answers AND keywords from the job description.
    * **Work Experience:** List the jobs the user provided. Place the transformed achievement bullet points under the correct job.
    * **Projects:** If any answers describe projects, create a dedicated "Projects" section.
    * **Education:** If the user mentioned their education, add it. If not, create a placeholder for them to fill in.

**Final Output Rules:**
* The tone must be professional, modern, and confident.
* The CV must be formatted as clean text.
* **Do not include any commentary, conversation, or text other than the CV document itself.**

---
**Target Job Description:**
%s

**User's Answers (JSON):**
%s
---

**Generated CV:**
`
