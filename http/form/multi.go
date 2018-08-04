package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

var html_header = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="content-type" content="text/html; charset=utf-8" />
		<title>Multipart Test</title>
	</head>
	<body>
<form enctype="multipart/form-data" action="/" method="post">
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" > 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交

	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="checkbox" name="dbnames" value="中文字符测试表单提交" checked> 中文字符测试表单提交
	<input type="file" name="file" >
	<input type="submit" />
</form>
</body></html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var f multipart.File
		fmt.Println(f)
		w.Write([]byte(" "))
		if http.MethodPost == r.Method {
			var err error
			f, _, err = r.FormFile("file")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

			w.Write([]byte(html_header)) //这里必须写在FormFile 之后
		} else {
			w.Write([]byte(html_header))
		}
		fmt.Println(f)

	})
	http.ListenAndServe(":12345", nil)
}
