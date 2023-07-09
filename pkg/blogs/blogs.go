package blogs

import "github.com/gin-gonic/gin"

//处理文章

// 定义文章方便存储数据
//文章的创建时间
// PostCreateTime string
//文章的最后编辑时间
// PostEditTime string
//文章的标题
// PostTitle string
//文章的标签
// PostTags string
//文章的id 根据这个筛选评论
// PostId int64
//文章原文 MarkDown
// PostText string
//渲染后的文章 html格式
// PostBody string

//提供三种查询文章的方式

func BlogIndex(c *gin.Context) {
	c.String(200, "ok")
}

// 根据时间查看文章 /yyyy-mm-dd /yyyy-mm /yyyy
func BlogDate(c *gin.Context) {
	dt := c.Param("date")
	c.String(200, dt)
}

// 根据id查看文章也是正确的打开文章方式
func BlogID(c *gin.Context) {
	bid := c.Param("id")
	c.String(200, bid)
}

// 根据标题打开文章，暂时隐藏不做使用
// func BlogTiele(c *gin.Context) {

// }
