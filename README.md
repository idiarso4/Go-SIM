# Academic Information System (SIA)

A robust academic information system built with Go, focusing on student assessment management and academic record keeping.

## 🚀 Features

### Assessment Management
- Create and manage student assessments
- Support for both Summative and Non-Summative evaluations
- Theory and Practical assessment categories
- Automated student list population by class
- Grade range validation (0-100)
- Optional remarks for each assessment

### Export Capabilities
- Export assessments to Excel format
- Standardized formatting with automatic styling
- Filename format: `nilai_[CLASS][SUBJECT][DATE].xlsx`
- Export includes:
  - Student list
  - Assessment details
  - Individual grades and remarks

### Search and Filter
- Date-based filtering
- Search by:
  - Class
  - Subject
  - Teacher
- Combined filter support

## 🛠 Technical Requirements

- Go 1.16 or higher
- SQLite database
- Modern web browser
- Internet connection

## 📦 Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/academic-system.git