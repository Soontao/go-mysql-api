package static

var SWAGGER_HTML = `
<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link href="https://cdn.bootcss.com/swagger-ui/3.0.20/swagger-ui.css" rel="stylesheet">
</head>

<body>

<div id="swagger-ui"></div>

<script src="https://cdn.bootcss.com/swagger-ui/3.0.20/swagger-ui-bundle.js"></script>
<script src="https://cdn.bootcss.com/swagger-ui/3.0.20/swagger-ui-standalone-preset.js"></script>
<script>
    window.onload = function () {

        // Build a system
        const ui = SwaggerUIBundle({
            url: "/api/swagger.json",
            dom_id: '#swagger-ui',
            deepLinking: true,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset
            ],
            plugins: [
                SwaggerUIBundle.plugins.DownloadUrl
            ],
            layout: "StandaloneLayout"
        })

        window.ui = ui
    }
</script>
</body>

</html>

`
