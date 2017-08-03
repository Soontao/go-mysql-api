package static

var SWAGGER_UI_HTML = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link href="https://cdn.bootcss.com/swagger-ui/3.1.2/swagger-ui.css" rel="stylesheet">
</head>
<style>
    body {
        margin: 0px !important;
    }
</style>

<body>

<div id="swagger-ui"></div>

<script src="https://cdn.bootcss.com/swagger-ui/3.1.2/swagger-ui-bundle.js"></script>
<script src="https://cdn.bootcss.com/swagger-ui/3.1.2/swagger-ui-standalone-preset.js"></script>
<script>
    window.onload = function () {

        function getSwaggerJSONURI() {
            l = window.location
            return l.protocol + "//" + l.host + "/api/swagger.json"
        }

        // Build a system
        const ui = SwaggerUIBundle({
            url: getSwaggerJSONURI(),
            dom_id: '#swagger-ui',
            deepLinking: true,
            filter: true,
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
