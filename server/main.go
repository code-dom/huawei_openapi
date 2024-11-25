package main

import (
	"io"
	"log"
	"net/http"
	"openapi/nlp"
	"openapi/ocr"
	tab "openapi/table"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// 图片上传并保存
func uploadImage(c *gin.Context) {
	// 获取上传的文件和文件头信息
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件上传失败"})
		return
	}
	defer file.Close()

	// 获取文件扩展名
	ext := filepath.Ext(fileHeader.Filename)
	// 使用时间戳生成唯一文件名
	fileName := time.Now().Format("20060102150405") + ext
	// 文件保存路径
	filePath := "./uploads/" + fileName

	// 创建文件夹用于存储上传的文件
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件夹创建失败"})
		return
	}

	// 创建文件并保存
	outFile, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败"})
		return
	}
	defer outFile.Close()

	// 将上传的文件内容复制到目标文件
	_, err = io.Copy(outFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	// 返回图片的访问URL
	imageURL := "http://localhost:8080/uploads/" + fileName
	absolutePath, err := filepath.Abs(filePath)
	content := ocr.OCR(absolutePath)
	c.JSON(http.StatusOK, gin.H{
		"message":  "图片上传成功",
		"imageUrl": imageURL,
		"content":  content,
	})
}

func translate(c *gin.Context) {
	// 从表单获取需要翻译的文本
	text := c.DefaultPostForm("text", "请输入需要翻译的文字")

	// 调用 NLP 服务进行翻译
	TranslatedText := nlp.GetNLP(text)

	if TranslatedText != "" {
		c.JSON(http.StatusOK, gin.H{
			"original":   text,
			"translated": TranslatedText,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "翻译失败",
		})
	}
}

func table(c *gin.Context) {
	// 获取上传的文件和文件头信息
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件上传失败"})
		return
	}
	defer file.Close()

	// 获取文件扩展名
	ext := filepath.Ext(fileHeader.Filename)
	// 使用时间戳生成唯一文件名
	fileName := time.Now().Format("20060102150405") + ext
	// 文件保存路径
	filePath := "./uploads/" + fileName
	// 创建文件夹用于存储上传的文件
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件夹创建失败"})
		return
	}
	// 创建文件并保存
	outFile, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败"})
		return
	}
	defer outFile.Close()
	// 将上传的文件内容复制到目标文件
	_, err = io.Copy(outFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		log.Fatal("获取文件绝对路径失败:", err)
	}
	content, rep := tab.Table(absolutePath)

	excelName := time.Now().Format("20060102150405") + ".xlsx"
	// 生成excel文件
	tab.GetExcel(rep, "./uploads/"+excelName)
	// 返回图片的访问URL
	excelUrl := "http://localhost:8080/uploads/" + excelName
	fileUrl := "http://localhost:8080/uploads/" + fileName

	c.JSON(http.StatusOK, gin.H{
		"message":  "上传成功",
		"imageUrl": fileUrl,
		"excelUrl": excelUrl,
		"text":     content,
	})
}

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 提供静态文件服务，允许访问/uploads目录
	r.Static("/uploads", "./uploads")

	r.LoadHTMLGlob("templates/*")

	// 定义上传接口，处理POST请求
	r.POST("/upload", uploadImage)
	r.POST("/translate", translate)
	r.POST("/table", table)

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "ocr.html", gin.H{})
	})

	r.GET("/ocr", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "ocr.html", gin.H{})
	})

	r.GET("/translate", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "nlp.html", gin.H{})
	})
	r.GET("/table", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "table.html", gin.H{})
	})

	// 启动HTTP服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
