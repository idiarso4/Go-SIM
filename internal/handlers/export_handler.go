package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"academic-system/internal/models"
	"academic-system/internal/repository"
)

func ExportAssessment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	assessment, err := assessmentRepo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assessment not found"})
		return
	}

	f := excelize.NewFile()
	sheet := "Sheet1"
	
    // Set headers with styling
    headers := []string{"No", "NIS", "Nama Siswa", "Nilai", "Keterangan", "Tanggal"}
    for i, header := range headers {
        col := string(rune('A' + i))
        f.SetCellValue(sheet, col+"1", header)
    }

    // Style headers
    headerStyle, _ := f.NewStyle(&excelize.Style{
        Font: &excelize.Font{
            Bold:   true,
            Size:   12,
            Color:  "#FFFFFF",
        },
        Fill: excelize.Fill{
            Type:    "pattern",
            Color:   []string{"#4472C4"},
            Pattern: 1,
        },
        Alignment: &excelize.Alignment{
            Horizontal: "center",
            Vertical:   "center",
        },
    })
    f.SetCellStyle(sheet, "A1", "F1", headerStyle)

    // Set column widths
    f.SetColWidth(sheet, "A", "A", 5)
    f.SetColWidth(sheet, "B", "B", 15)
    f.SetColWidth(sheet, "C", "C", 30)
    f.SetColWidth(sheet, "D", "D", 10)
    f.SetColWidth(sheet, "E", "E", 20)
    f.SetColWidth(sheet, "F", "F", 15)

	// Fill data
	for i, student := range assessment.Students {
		row := i + 2
        rowStr := strconv.Itoa(row)
        
        f.SetCellValue(sheet, "A"+rowStr, i+1)
        f.SetCellValue(sheet, "B"+rowStr, student.NIS)
        f.SetCellValue(sheet, "C"+rowStr, student.Name)
        f.SetCellValue(sheet, "D"+rowStr, student.Score)
        f.SetCellValue(sheet, "E"+rowStr, getRemarks(student.Score))
        f.SetCellValue(sheet, "F"+rowStr, time.Now().Format("2006-01-02"))
	}

    // Create exports directory if it doesn't exist
    exportDir := "./exports"
    if err := os.MkdirAll(exportDir, 0755); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create export directory"})
        return
    }

	// Generate filename
    filename := fmt.Sprintf("%s/nilai_%s_%s_%s.xlsx",
        exportDir,
        assessment.Class.Name,
        assessment.Subject.Name,
        time.Now().Format("20060102_150405"))

	if err := f.SaveAs(filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.File(filename)
}

func getRemarks(score float64) string {
    switch {
    case score >= 90:
        return "Sangat Baik"
    case score >= 80:
        return "Baik"
    case score >= 70:
        return "Cukup"
    case score >= 60:
        return "Kurang"
    default:
        return "Perlu Perbaikan"
    }
}