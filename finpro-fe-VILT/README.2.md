# 🧠 Flow Penghitungan Skor → AI Recommendation (Gemini)

Aku buatkan flow yang sesuai dengan sistem Lumora kamu sekarang:

* Self-assessment
* Scoring
* Kategorisasi
* Gemini recommendation
* Simpan ke database

---

# 📊 1️⃣ USER MENGISI SELF-ASSESSMENT

User menjawab 20 pertanyaan:

| Category        | Question |
| --------------- | -------- |
| Planning        | Q1–Q5    |
| Time Management | Q6–Q10   |
| Cognitive       | Q11–Q15  |
| Reflection      | Q16–Q20  |

Jawaban menggunakan skala:

```
1 = Strongly Disagree
2 = Disagree
3 = Neutral
4 = Agree
5 = Strongly Agree
```

---

# 🗄 2️⃣ SYSTEM MENYIMPAN JAWABAN

Disimpan ke:

```
assessment_responses
```

Contoh:

| user_id | question_id | answer_value |
| ------- | ----------- | ------------ |
| 1       | 1           | 4            |
| 1       | 2           | 5            |

---

# 🧮 3️⃣ SYSTEM MENGHITUNG SKOR

## Formula

```
Planning Score
= Q1 + Q2 + Q3 + Q4 + Q5
```

```
Time Management Score
= Q6 + Q7 + Q8 + Q9 + Q10
```

```
Cognitive Score
= Q11 + Q12 + Q13 + Q14 + Q15
```

```
Reflection Score
= Q16 + Q17 + Q18 + Q19 + Q20
```

---

# 📊 4️⃣ SYSTEM MELAKUKAN KATEGORISASI

## Rule

| Score | Category |
| ----- | -------- |
| 5–12  | Low      |
| 13–18 | Medium   |
| 19–25 | High     |

---

## Contoh hasil

| Dimension       | Score | Category |
| --------------- | ----- | -------- |
| Planning        | 10    | Low      |
| Time Management | 16    | Medium   |
| Cognitive       | 22    | High     |
| Reflection      | 11    | Low      |

---

# 🗄 5️⃣ HASIL DISIMPAN

Disimpan ke:

```
result_summary
```

Contoh:

```
{
  "planning_score": 10,
  "time_management_score": 16,
  "cognitive_score": 22,
  "reflection_score": 11,
  "category_result": "mixed"
}
```

---

# 🤖 6️⃣ SYSTEM MEMBENTUK PROMPT GEMINI

Backend membuat prompt berdasarkan hasil kategori.

---

## Contoh Prompt

```
User learning profile:

Planning: Low
Time Management: Medium
Cognitive Strategy: High
Reflection: Low

Generate:
1. learning analysis
2. strengths
3. weaknesses
4. personalized study recommendations
5. practical actions for improvement

Response should be concise, motivational, and suitable for students.
```

---

# 🚀 7️⃣ BACKEND MENGIRIM KE GEMINI API

Flow:

```text id="53ptc4"
Backend
↓
HTTP Request
↓
Gemini API
↓
AI Response
```

---

## Contoh response Gemini

```
Strength:
You have strong cognitive learning ability and understand materials well.

Weakness:
Your learning planning and reflection habits are still inconsistent.

Recommendation:
- Create a fixed weekly study schedule
- Use daily reflection notes after studying
- Break large tasks into smaller goals
```

---

# 🗄 8️⃣ HASIL AI DISIMPAN

Ke tabel:

```
recommendation
```

---

