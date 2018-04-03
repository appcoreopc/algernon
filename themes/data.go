package themes

var (
	// builtinCodeStyles is a map of the themes names corresponding to chroma styles
	// See the Chroma Style Gallery for more styles: https://xyproto.github.io/splash/docs/
	builtinCodeStyles = map[string]string{"gray": "github", "dark": "monokai", "redbox": "fruity", "bw": "github", "wing": "github", "material": "monokailight", "neon": "tango"}

	// Themes is a map over the available built-in CSS themes. Corresponds with the font themes below.
	// "default" and "gray" are equal. "default" should never be used directly, but is here as a safeguard.
	builtinThemes = map[string][]byte{
		"default":  []byte("@import url(//fonts.googleapis.com/css?family=Lato:300); body { background-color: #e7eaed; color: #0b0b0b; font-family: 'Lato', sans-serif; font-weight: 300;  margin: 4.5em; font-size: 1em; } a { color: #401010; font-family: courier; } a:hover { color: #801010; } a:active { color: yellow; } h1 { color: #101010; } img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #ccc; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){background-color: #efefef; } table tr:nth-child(odd) { background-color: #fff; } table tr:hover { background-color: #e0f0ff; color: black; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #666; color: white; } code { background-color: rgba(50, 50, 50, 0.1); color: #202020; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; font-size: 90%; word-wrap: break-word; font-family: Menlo, Monaco, Consolas, \"Courier New\", monospace; }"),
		"gray":     []byte("@import url(//fonts.googleapis.com/css?family=Lato:300); body { background-color: #e7eaed; color: #0b0b0b; font-family: 'Lato', sans-serif; font-weight: 300;  margin: 4.5em; font-size: 1em; } a { color: #401010; font-family: courier; } a:hover { color: #801010; } a:active { color: yellow; } h1 { color: #101010; } img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #ccc; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){background-color: #efefef; } table tr:nth-child(odd) { background-color: #fff; } table tr:hover { background-color: #e0f0ff; color: black; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #666; color: white; } code { background-color: rgba(50, 50, 50, 0.1); color: #202020; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; font-size: 90%; word-wrap: break-word; font-family: Menlo, Monaco, Consolas, \"Courier New\", monospace; }"),
		"light":    []byte("@import url(//fonts.googleapis.com/css?family=Lato:300); body { background-color: #eaeaea; color: #0b0b0b; font-family: 'Lato', sans-serif; font-weight: 300;  margin: 4.5em; font-size: 1em; } a { color: #401010; font-family: courier; } a:hover { color: #801010; } a:active { color: yellow; } h1 { color: #101010; } img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #ccc; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){background-color: #efefef; } table tr:nth-child(odd) { background-color: #fff; } table tr:hover { background-color: #e0f0ff; color: black; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #666; color: white; } code { background-color: rgba(50, 50, 50, 0.1); color: #202020; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; font-size: 90%; word-wrap: break-word; font-family: Menlo, Monaco, Consolas, \"Courier New\", monospace; }"),
		"dark":     []byte("@import url(//fonts.googleapis.com/css?family=Lato:400); body { background-color: #101010; color: #f0f0f0; font-family: 'Lato', sans-serif; font-weight: 400;  margin: 4.5em; font-size: 1em; } a { color: #c0a0a0; font-family: courier; } a:hover { color: #f0a0a0; } a:active { color: yellow; } h1 { color: #f0f0f0; } img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #999; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){background-color: #222; } table tr:nth-child(odd) { background-color: #333; } table tr:hover { background-color: #333; color: #ff8060; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #d7d7d7; color: #292929; } code { background-color: rgba(255, 255, 255, 0.18); color: #e0e0e0; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; font-size: 90%; word-wrap: break-word; font-family: Menlo, Monaco, Consolas, \"Courier New\", monospace; }"),
		"redbox":   []byte("@import url(//fonts.googleapis.com/css?family=Monoton|Monofett);html{background-color:#222;}body{color:#111;background-color:#999;font-family:Impact,'Arial Black',sans-serif;font-size:1.7em;margin:2.7em;padding:0 5em 1em 2em;border-radius:50px;border:solid 10px #a00;box-shadow:10px 10px 16px black, 6px 6px 8px #222 inset;}h2{font-family:Monoton,cursive;color:black;}ul{margin-left:1em;}a{text-decoration:none;color:#b00;}a:hover{color:#dc0;}code{font-family:Monofett,cursive;} img { max-width: 100%; } code { background-color: rgba(255, 255, 255, 0.18); color: #e0e0e0; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; font-size: 90%; word-wrap: break-word; font-family: Menlo, Monaco, Consolas, \"Courier New\", monospace; }"),
		"bw":       []byte("body { font-family: BlinkMacSystemFont, \"Segoe UI\", Helvetica, sans-serif; margin: 1px; padding: 1px; } button { background: black; border-radius: 4px; color: #ffffff; font-size: 20px; padding: 10px 20px 10px 20px; border: none; outline: none; margin-left: 2px; margin-right: 2px; } button:hover { background: orange; text-decoration: none; } button:disabled { background: gray; color: #eee; } input { font-size: 16px; padding: 10px; border: 1px solid #999; } a { padding: 1px; text-decoration: none; color: black; } a:hover { padding: 1px; border-bottom: 1px solid black; } a.button { font-size: 1em; border: 1px solid black; background-color: white; color: black; width: 150px; padding: 10px 0px; margin-right: 10px; margin-top: 5px; margin-bottom: 5px; text-align: center; display: inline-block; } a.button:hover { text-decoration: none; color: white; background-color: black;   } a.button:active { color: white; background-color: black;  } a.go-back { position: absolute; right: 10px; top: 10px; padding: 10px; font-size: 0.8em; background-color: black; color: white; } .ha { box-sizing: border-box; margin: 0 150px; padding: 20px; line-height: 1.5; font-size: 15.5px; line-height: 1.55; } .ha ul { list-style: none; padding: 0 !important; column-count: 2; max-width: 0; column-gap: 20px; min-width: 350px; } .ha ul li { margin-left: 0px; } .ha h1 { margin-top: 0; margin-bottom: 0; } .ha h1 { font-size: 32px; font-weight: 200; margin-top: 0px; margin-bottom: 16px; } .ha h5 { margin-bottom: 2px; } .ha hr { padding: 0; margin: 24px 0; background-color: #e7e7e7; border: 0; height: 0.09em; } .ha li { margin-top: 0.25em; } .gh-btns { float: right; margin-top: -4px; }"),                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // from MIT licensed HyperApp: hyperapp.glitch.me/style.css
		"wing":     []byte(`body,html{margin:0;padding:0}html{box-sizing:border-box;font-size:62.5%}body{line-height:1.6;font-size:1.5rem;font-weight:400;font-family:-apple-system,BlinkMacSystemFont,Avenir,"Avenir Next","Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Fira Sans","Droid Sans","Helvetica Neue",sans-serif}h1,h2,h3,h4,h5,h6{margin-top:0;margin-bottom:2rem;font-weight:400;font-family:-apple-system,BlinkMacSystemFont,Avenir,"Avenir Next","Segoe UI",Roboto,Oxygen,Ubuntu,Cantarell,"Fira Sans","Droid Sans","Helvetica Neue",sans-serif}h1{font-size:5rem;line-height:1.2}h2{font-size:4.2rem;line-height:1.25}h3{font-size:3.6rem;line-height:1.3}h4{font-size:3rem;line-height:1.35}h5{font-size:2.4rem;line-height:1.5}h6{font-size:1.8rem;line-height:1.6}p{margin-top:0;margin-bottom:2rem;font-size:1.5rem}@media (max-width:400px){h1{font-size:4rem}h2{font-size:3.5rem}h3{font-size:3rem}h4{font-size:2.6rem}h5{font-size:2.2rem}h6{font-size:1.8rem}}a{color:#104cfb;transition:color .1s ease}a:hover{cursor:pointer;color:#222}.button,[type=submit],button{padding:1.1rem 3.5rem;margin-top:0;margin-bottom:2rem;color:#f5f5f5;font-size:1.5rem;background:#111;border:1.5px solid #111;border-radius:2px;transition:all .2s ease}.button.outline,[type=submit].outline,button.outline{color:#111;background:0 0;border:1.5px solid #111}.button:hover,[type=submit]:hover,button:hover{color:#f5f5f5;background:#222}.button.outline:hover,[type=submit].outline:hover,button.outline:hover{background:0 0;border:1.5px solid #222;color:#222}.button:focus,[type=submit]:focus,button:focus{outline:0}.button:active,[type=submit]:active,button:active{transform:scale(.99)}.button.disabled,.button:disabled,[type=submit]:disabled,button:disabled{opacity:.8;cursor:not-allowed}input[type=email],input[type=file],input[type=number],input[type=password],input[type=search],input[type=tel],input[type=text],input[type=url],select,textarea,textarea[type=text]{width:100%;height:45px;padding:10px 10px;margin-top:1rem;margin-bottom:1rem;font-size:1.3rem;background:#f1f1f1;border:1px solid #a4a4a4;border-radius:2px;box-sizing:border-box;transition:all .2s ease}input[type=email]:hover,input[type=file]:hover,input[type=number]:hover,input[type=password]:hover,input[type=search]:hover,input[type=tel]:hover,input[type=text]:hover,input[type=url]:hover,select:hover,textarea:hover,textarea[type=text]:hover{border:1px solid #111}input[type=email]:focus,input[type=file]:focus,input[type=number]:focus,input[type=password]:focus,input[type=search]:focus,input[type=tel]:focus,input[type=text]:focus,input[type=url]:focus,select:focus,textarea:focus,textarea[type=text]:focus{outline:0;border:1px solid #104cfb}textarea,textarea[type=text]{min-height:7rem}.container{max-width:960px;margin:0 auto;width:80%}.row{display:flex;flex-flow:row wrap;justify-content:space-between;width:100%;margin:1rem 0}.row>:first-child{margin-left:0}.col{flex:1 1 0px}.col,[class*=" col-"],[class^=col-]{margin-left:4%}.col-1{width:4.333333333333332%}.col-2{width:12.666666666666664%}.col-3{width:21%}.col-4{width:29.33333333333333%}.col-5{width:37.66666666666667%}.col-6{width:46%}.col-7{width:54.333333333333336%}.col-8{width:62.66666666666666%}.col-9{width:71%}.col-10{width:79.33333333333334%}.col-11{width:87.66666666666666%}.col-12{width:96%}@media screen and (max-width:768px){.col,[class*=" col-"],[class^=col-]{margin:1rem 0;flex:0 0 100%}}ol,ul{padding-left:0;margin-top:0;margin-bottom:2rem;list-style-position:inside}ol ol,ol ul,ul ol,ul ul{margin:1rem 0 1rem 2rem;font-size:95%}li{margin-bottom:1rem}.table{width:100%;border:none;border-collapse:collapse;border-spacing:0;text-align:left}.table td,.table th{vertical-align:middle;padding:12px 4px}.table thead{border-bottom:2px solid #333030}@media screen and (max-width:768px){.table.responsive{position:relative;display:block}.table.responsive td,.table.responsive th{margin:0}.table.responsive thead{display:block;float:left;border:0}.table.responsive thead tr{display:block;padding:0 10px 0 0;border-right:2px solid #333030}.table.responsive th{display:block;text-align:right}.table.responsive tbody{display:block;overflow-x:auto;white-space:nowrap}.table.responsive tbody tr{display:inline-block}.table.responsive td{display:block;min-height:16px;text-align:left}.table.responsive tr{padding:0 10px}}img{margin-top:0;margin-bottom:2rem}.nav{position:relative;display:flex;flex-wrap:wrap;align-items:center;padding-top:1rem;padding-bottom:1rem}.nav-item{padding:1rem 2rem}.cards{display:flex;flex-wrap:wrap;justify-content:space-between;overflow:hidden;margin-bottom:2rem}.card{flex-direction:column;flex:0 1 calc(50% - .5rem);box-shadow:0 5px 15px rgba(0,0,0,.1);margin-bottom:2rem}.card-image{display:block;position:relative}.card-image img{display:block;height:auto;width:100%}.card-header{font-weight:600;margin:0;padding:2rem 3rem 1rem}.card-body{padding:0 3rem 2rem 3rem;min-height:100px}.card-footer{display:flex;align-items:stretch;border-top:1px solid #dbdbdb;flex:1}.card-footer .card-footer-item{display:flex;flex-basis:0;flex-grow:1;flex-shrink:0;align-items:center;justify-content:center;margin:0;padding:1rem}.card-footer-item:not(:first-child){border-left:1px solid #dbdbdb}@media (max-width:768px){.cards>.card{flex:auto}}.pull-right{float:right}.pull-left{float:left}.text-center{text-align:center}.text-left{text-align:left}.text-right{text-align:right}.full-screen{width:100%;min-height:100vh}.full-width{width:100%}.vertical-align{display:flex;align-items:center}.horizontal-align{display:flex;justify-content:center}.center{display:flex;align-items:center;justify-content:center;flex-direction:column}.right{display:flex;align-items:center;justify-content:flex-end}.left{display:flex;align-items:center;justify-content:flex-start}.fixed{position:fixed;width:100%}@media screen and (max-width:400px){.hide-phone{display:none}}@media screen and (max-width:768px){.hide-tablet{display:none}}pre{margin-top:0;margin-bottom:2rem}code{padding:.2rem .5rem;margin:0 .2rem;font-size:1.3rem;white-space:nowrap;background:#f1f1f1;border:1px solid #dbdbdb;border-radius:4px;font-family:Consolas,Monaco,Menlo,monospace}pre>code{display:block;padding:1rem 1.5rem;white-space:pre-wrap;word-wrap:break-word}`), // From https://github.com/kbrsh/wing/blob/master/dist/wing.min.css
		"material": []byte("body { margin: 3em; } pre{margin-top:0;margin-bottom:2rem}code{padding:.2rem .5rem;margin:0 .2rem;font-size:1.3rem;white-space:nowrap;background:#f1f1f1;border:1px solid #dbdbdb;border-radius:4px;font-family:Consolas,Monaco,Menlo,monospace}pre>code{display:block;padding:1rem 1.5rem;white-space:pre-wrap;word-wrap:break-word} img { max-width: 100%; } table { font-family: Helvetica, sans-serif; border-collapse: collapse; width: 100%; font-weight: bold; } table td, table th { border: 1px solid #ccc; padding: 0.45em 0.55em 0.45em 0.55em; } table tr:nth-child(even){background-color: #efefef; } table tr:nth-child(odd) { background-color: #fff; } table tr:hover { background-color: #e0f0ff; color: black; } table th { padding-top: 12px; padding-bottom: 12px; text-align: left; background-color: #666; color: white; } code { background-color: rgba(50, 50, 50, 0.1); color: #202020; border-radius: 0.3em; padding: 0.1em 0.4em 0.1em 0.4em; font-size: 90%; word-wrap: break-word; font-family: Menlo, Monaco, Consolas, \"Courier New\", monospace; }"),                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       // CSS and JS is added with a special function
		"neon":     []byte("html{line-height:1}body{align-items:center;background-color:#111;display:flex;font-family:Helvetica Neue,sans-serif;height:90vh;justify-content:center;margin:0;padding:0;text-align:center}h1{color:#00caff;font-weight:100;font-size:8em;margin:0;padding-bottom:15px}button{background:#111;border:1px solid #00caff;color:#00caff;font-size:2em;font-weight:100;margin:0;outline:none;padding:5px 15px;transition:background .2s;&:hover,&:active,&:disabled{background:#00caff;color:#111}&:active{outline:2px solid #00caff}&:focus{border:1px solid #00caff}}button + button{margin-left:3px}pre{margin-top:0;margin-bottom:2rem}code{padding:.2rem .5rem;margin:0 .2rem;font-size:1.3rem;white-space:nowrap;background:#f1f1f1;border:1px solid #dbdbdb;border-radius:4px;font-family:Consolas,Monaco,Menlo,monospace}pre>code{display:block;padding:1rem 1.5rem;white-space:pre-wrap;word-wrap:break-word}*{box-sizing:border-box;margin:0;padding:0}body{background:#111}.container{align-items:center;display:flex;height:100vh;justify-content:center;text-align:center}input{background:rgba(0,0,0,.2);color:#00caff;border:1px solid #111;font:4em Helvetica Neue,sans-serif;font-weight:300;height:100vh;width:100vw;outline:0;padding:0 40px;position:absolute;&:hover{background:rgba(0,0,0,0)}}img{width:90vw;height:90vh}"),
	}
)
