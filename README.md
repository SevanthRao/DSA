<body class=" "><div id="MathJax_Message" style="display: none;"></div><div style="visibility: hidden; overflow: hidden; position: absolute; top: 0px; height: 1px; width: auto; padding: 0px; border: 0px; margin: 0px; text-align: left; text-indent: 0px; text-transform: none; line-height: normal; letter-spacing: normal; word-spacing: normal;"><div id="MathJax_Hidden"></div></div><span style="display:none;" class="csrf-token" data-csrf="c2dbc165e02d03cca30cbbb627118425">&nbsp;</span>

<!-- .notificationTextCleaner used in Codeforces.showAnnouncements() -->
<div class="notificationTextCleaner" style="font-size: 0"></div>
<div class="button-up" style="display: none; opacity: 0.7; width: 50px; height:100%; position: fixed; left: 0; top: 0; cursor: pointer; text-align: center; line-height: 35px; color: #d3dbe4; font-weight: bold; font-size: 3.0rem;"><i class="icon-circle-arrow-up"></i></div>
<div class="verdictPrototypeDiv" style="display: none;"></div>

<!-- Codeforces JavaScripts. -->
<script type="text/javascript">
    String.prototype.hashCode = function() {
        var hash = 0, i, chr;
        if (this.length === 0) return hash;
        for (i = 0; i < this.length; i++) {
            chr   = this.charCodeAt(i);
            hash  = ((hash << 5) - hash) + chr;
            hash |= 0; // Convert to 32bit integer
        }
        return hash;
    };

    var queryMobile = Codeforces.queryString.mobile;
    if (queryMobile === "true" || queryMobile === "false") {
        Codeforces.putToStorage("useMobile", queryMobile === "true");
    } else {
        var useMobile = Codeforces.getFromStorage("useMobile");
        if (useMobile === true || useMobile === false) {
            if (useMobile != false) {
                Codeforces.redirect(Codeforces.updateUrlParameter(document.location.href, "mobile", useMobile));
            }
        }
    }
</script>

<script type="text/javascript">
    if (window.parent.frames.length > 0) {
        window.stop();
    }
</script>





    <script type="text/javascript">
        $(document).ready(function () {
    (function () {
        jQuery.expr[':'].containsCI = function(elem, index, match) {
            return !match || !match.length || match.length < 4 || !match[3] || (
                    elem.textContent || elem.innerText || jQuery(elem).text() || ''
            ).toLowerCase().indexOf(match[3].toLowerCase()) >= 0;
        }
    }(jQuery));

    $.ajaxPrefilter(function(options, originalOptions, xhr) {
        var csrf = Codeforces.getCsrfToken();

        if (csrf) {
            var data = originalOptions.data;
            if (originalOptions.data !== undefined) {
                if (Object.prototype.toString.call(originalOptions.data) === '[object String]') {
                    data = $.deparam(originalOptions.data);
                }
            } else {
                data = {};
            }
            options.data = $.param($.extend(data, { csrf_token: csrf }));
        }
    });

    window.getCodeforcesServerTime = function(callback) {
        $.post("/data/time", {}, callback, "json");
    }

    window.updateTypography = function () {
        $("div.ttypography code").addClass("tt");
        $("div.ttypography pre>code").addClass("prettyprint").removeClass("tt");
        $("div.ttypography table").addClass("bordertable");
        prettyPrint();
    }

    $.ajaxSetup({ scriptCharset: "utf-8" ,contentType: "application/x-www-form-urlencoded; charset=UTF-8", headers: {
        'X-Csrf-Token': Codeforces.getCsrfToken()
    }, beforeSend: function(jqXHR, settings) {
        const url = new URL(settings.url, window.location.href);
        const urlHref = url.href;
        url.searchParams.set('rv', Math.random().toString(36).substr(2, 9)); // Generate a random value
        settings.url = url.href;
    }});

    window.updateTypography();
    Codeforces.signForms();
    Codeforces.applyDatatableJs("Hold Ctrl+Alt and click to hide column. Also, you can drag&drop columns to reorder them.");

    setTimeout(function() {
        $(".second-level-menu-list").lavaLamp({
            fx: "backout",
            speed: 700
        });
    }, 100);


    Codeforces.countdown();
    $("a[rel='photobox']").colorbox();

        var count = 0;

        function getDelay() {
            var delay = 0;
            var last = Codeforces.getFromStorage("LastOnlineTimeUpdaterMillis", 0);
            if (last && last != null) {
                var period = count < 20 ? 360 * 1000 : 600 * 1000; // 6 or 10 minutes.
                delay = period - (new Date().getTime() - last);
                if (delay < 0)
                    delay = 0;
            }
            return delay;
        }

        window.setInterval(function () {
            if (getDelay() <= 0 && count < 120) {
                ++count;
                Codeforces.ping("/data/update-online");
            }
        }, 5000);

        var handle = "sevanthrxo";
        $("a.rated-user:contains(" + handle + "), span.participant:contains(" + handle + ")").each(function () {
            if ($(this).text() == handle) {
                var td = $(this).parent();
                var tr = $(this).parent().parent();
                if (td.is("td") && tr.is("tr") && !tr.hasClass("disable-highlight-self")) {
                    info(handle);
                    tr.addClass("highlighted-row");
                }
            }
        });

    function showAnnouncements(json) {
        if (json.t === "a") {
            setTimeout(function() {
               Codeforces.showAnnouncements(json.d, "en");
            }, Math.random() * 500);
        }
    }

    function showEventCatcherUserMessage(json) {
        if (json.t === "s") {
            var points = json.d[5];
            var passedTestCount = json.d[7];
            var judgedTestCount = json.d[8];
            var verdict = preparedVerdictFormats[json.d[12]];
            var verdictPrototypeDiv = $(".verdictPrototypeDiv");
            verdictPrototypeDiv.html(verdict);
            if (judgedTestCount != null && judgedTestCount != undefined) {
                verdictPrototypeDiv.find(".verdict-format-judged").text(judgedTestCount);
            }
            if (passedTestCount != null && passedTestCount != undefined) {
                verdictPrototypeDiv.find(".verdict-format-passed").text(passedTestCount);
            }
            if (points != null && points != undefined) {
                verdictPrototypeDiv.find(".verdict-format-points").text(points);
            }
            Codeforces.showMessage(verdictPrototypeDiv.text());
        }
    }

    $(".clickable-title").each(function() {
        var title = $(this).attr("data-title");
        if (title) {
            var tmp = document.createElement("DIV");
            tmp.innerHTML = title;
            $(this).attr("title", tmp.textContent || tmp.innerText || "");
        }
    });

    $(".clickable-title").click(function() {
        const title = $(this).attr("data-title");
        const clazz = $(this).attr("data-clazz");
        const props = {};
        if (typeof (clazz) !== 'undefined') {
            props['clazz'] = clazz;
        }
        if (title) {
            Codeforces.alert(title, null, null, props);
        } else {
            Codeforces.alert($(this).attr("title"), null, null, props);
        }
    }).css("position", "relative").css("bottom", "3px");

        Codeforces.showDelayedMessage();

    Codeforces.reformatTimes();

    //Codeforces.initializePubSub();
    if (window.codeforcesOptions.subscribeServerUrl) {
        window.eventCatcher = new EventCatcher(
            window.codeforcesOptions.subscribeServerUrl,
            [
                Codeforces.getGlobalChannel(),
                Codeforces.getUserChannel(),
                Codeforces.getUserShowMessageChannel(),
                Codeforces.getContestChannel(),
                Codeforces.getParticipantChannel(),
                Codeforces.getTalkChannel()
            ]
        );

        if (Codeforces.getParticipantChannel()) {
            window.eventCatcher.subscribe(Codeforces.getParticipantChannel(), function(json) {
                showAnnouncements(json);
            });
        }

        if (Codeforces.getContestChannel()) {
            window.eventCatcher.subscribe(Codeforces.getContestChannel(), function(json) {
                showAnnouncements(json);
            });
        }

        if (Codeforces.getGlobalChannel()) {
            window.eventCatcher.subscribe(Codeforces.getGlobalChannel(), function(json) {
                showAnnouncements(json);
            });
        }

        if (Codeforces.getUserChannel()) {
            window.eventCatcher.subscribe(Codeforces.getUserChannel(), function(json) {
                showAnnouncements(json);
            });
        }

        if (Codeforces.getUserShowMessageChannel()) {
            window.eventCatcher.subscribe(Codeforces.getUserShowMessageChannel(), function(json) {
                showEventCatcherUserMessage(json);
            });
        }
    }

    Codeforces.setupContestTimes("/group/4vcXCPx8NY/data/contests");
    Codeforces.setupSpoilers();
    Codeforces.setupTutorials("/data/problemTutorial");
        });
    </script>


<div id="body">
<div style="width: 950px; margin: 0 auto;" class="compact-problemset">
    <div id="header" style="position: relative; margin: 0;">
        <div style="float:left;">
                    <a href="/"><img height="65" style="height: 65px;" src="//codeforces.com/codeforces.org/s/19965/images/codeforces-sponsored-by-ton.png" alt="Codeforces"></a>
        </div>
        <div class="lang">
            <div style="text-align: right;">
                <a href="?locale=en"><img src="//codeforces.com/codeforces.org/s/19965/images/flags/24/gb.png" title="In English" alt="In English"></a>
                <a href="?locale=ru"><img src="//codeforces.com/codeforces.org/s/19965/images/flags/24/ru.png" title="По-русски" alt="По-русски"></a>
            </div>
        </div>
        <br style="clear: both;">
    </div>

    <br style="clear: both;">

    <div style="text-align: center; font-size: 1.8rem; margin-bottom: 0.5em;" class="caption">Introduction to C++</div>
    <div style="border-top: 1px solid #ccc; height: 1em;"></div>

    <div class="problem-frames">
        <div style="margin-bottom: 2em;">

    <style>
        #facebox .content:has(.diff-popup) {
            width: 90vw;
            max-width: 120rem !important;
        }

        .testCaseMarker {
            position: absolute;
            font-weight: bold;
            font-size: 1rem;
        }

        .diff-popup {
            width: 90vw;
            max-width: 120rem !important;
            display: none;
            overflow: auto;
        }

        .input-output-copier {
            font-size: 1.2rem;
            float: right;
            color: #888 !important;
            cursor: pointer;
            border: 1px solid rgb(185, 185, 185);
            padding: 3px;
            margin: 1px;
            line-height: 1.1rem;
            text-transform: none;
        }

        .input-output-copier:hover {
            background-color: #def;
        }

        .test-explanation textarea {
            width: 100%;
            height: 1.5em;
        }

        .pending-submission-message {
            color: darkorange !important;
        }
    </style>
    <script>
        const OPENING_SPACE = String.fromCharCode(1001);
        const CLOSING_SPACE = String.fromCharCode(1002);

        const nodeToText = function (node, pre) {
            let result = [];

            if (node.tagName === "SCRIPT" || node.tagName === "math"
                || (node.classList && node.classList.contains("input-output-copier")))
                return [];

            if (node.tagName === "NOBR") {
                result.push(OPENING_SPACE);
            }

            if (node.nodeType === Node.TEXT_NODE) {
                let s = node.textContent;
                if (!pre) {
                    s = s.replace(/\s+/g, " ");
                }
                if (s.length > 0) {
                    result.push(s);
                }
            }

            if (pre && node.tagName === "BR") {
                result.push("\n");
            }

            node.childNodes.forEach(function (child) {
                result.push(nodeToText(child, node.tagName === "PRE").join(""));
            });

            if (node.tagName === "DIV"
                || node.tagName === "P"
                || node.tagName === "PRE"
                || node.tagName === "UL"
                || node.tagName === "LI"
            ) {
                result.push("\n");
            }

            if (node.tagName === "NOBR") {
                result.push(CLOSING_SPACE);
            }

            return result;
        }

        const isSpecial = function (c) {
            return c === ',' || c === '.' || c === ';' || c === ')' || c === ' ';
        }

        const convertStatementToText = function(statmentNode) {
            const text = nodeToText(statmentNode, false).join("").replace(/ +/g, " ").replace(/\n\n+/g, "\n\n");
            let result = [];
            for (let i = 0; i < text.length; i++) {
                const c = text.charAt(i);
                if (c === OPENING_SPACE) {
                    if (!((i > 0 && text.charAt(i - 1) === '(') || (result.length > 0 && result[result.length - 1] === ' '))) {
                        result.push('+');
                    }
                } else if (c === CLOSING_SPACE) {
                    if (!(i + 1 < text.length && isSpecial(text.charAt(i + 1)))) {
                        result.push('-');
                    }
                } else {
                    result.push(c);
                }
            }
            return result.join("").split("\n").map(value => value.trim()).join("\n");
        };
    </script>
    <div class="diff-popup">
    </div>

<div class="problemindexholder" problemindex="A" data-uuid="ps_2e9f0ef841f80fa0003f83614b16050006cb4ab7">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">A. Hello World</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>Your task is to print the string <span class="tex-font-style-tt">Hello World</span>.</p></div><div class="input-specification"><div class="section-title">Input</div><p>There is no input for this problem.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print exactly: </p><pre class="verbatim"><br>Hello World<br></pre></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id005825899579715637" id="id009757059177815821" class="input-output-copier">Copy</div></div><pre id="id005825899579715637">There is no input.
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id0040310859814294453" id="id008500093912875626" class="input-output-copier">Copy</div></div><pre id="id0040310859814294453">Hello World</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {
            window.changedTests = new Set();

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");



        String.prototype.replaceAll = function (search, replace) {
            return this.split(search).join(replace);
        };

        $(".sample-test .title").each(function () {
            const preId = ("id" + Math.random()).replaceAll(".", "0");
            const cpyId = ("id" + Math.random()).replaceAll(".", "0");

            $(this).parent().find("pre").attr("id", preId);
            const $copy = $("<div title='Copy' data-clipboard-target='#" + preId + "' id='" + cpyId + "' class='input-output-copier'>Copy</div>");
            $(this).append($copy);

            const clipboard = new Clipboard('#' + cpyId, {
                text: function (trigger) {
                    const pre = document.querySelector('#' + preId);
                    const lines = pre.querySelectorAll(".test-example-line");
                    return Codeforces.filterClipboardText(pre.innerText, lines.length);
                }
            });

            const isInput = $(this).parent().hasClass("input");

            clipboard.on('success', function (e) {
                if (isInput) {
                    Codeforces.showMessage("The example input has been copied into the clipboard");
                } else {
                    Codeforces.showMessage("The example output has been copied into the clipboard");
                }
                e.clearSelection();
            });
        });

        $(".test-form-item input").change(function () {
            addPendingSubmissionMessage($($(this).closest("form")), "You changed the answer, do not forget to submit it if you want to save the changes");
            const index = $(this).closest(".problemindexholder").attr("problemindex");
            let test = "";
            $(this).closest("form input").each(function () {
                const test_ = $(this).attr("name");
                if (test_ && test_.substring(0, 4) === "test") {
                    test = test_;
                }
            });
            if (index.length > 0 && test.length > 0) {
                const indexTest = index + "::" + test;
                window.changedTests.add(indexTest);
            }
        });

        $(window).on('beforeunload', function () {
            if (window.changedTests.size > 0) {
                return 'Dialog text here';
            }
        });

        autosize($('.test-explanation textarea'));

        function showTestCaseMarker($input, tp, clazz, end) {
            let top = 1E20;
            let left = 1E20;
            let problem = $input.closest(".problemindexholder");
            $input.find("." + clazz).css("background-color", "#FFFDE7").each(function() {
                top = Math.min(top, $(this).offset().top);
                left = Math.min(left, $(this).offset().left);
            });
            let testCaseMarker = problem.find(".testCaseMarker_" + tp + "_" + end);
            if (testCaseMarker.length === 0) {
                const html = "<div class=\"testCaseMarker testCaseMarker_" + tp + "_" + end + " notice\"></div>";
                problem.append($(html));
                testCaseMarker = problem.find(".testCaseMarker_" + tp + "_" + end);
            }
            if (testCaseMarker) {
                testCaseMarker.show()
                    .offset({top, left: left - 20})
                    .text(end);
            }
        }

        $(".test-example-line").hover(function() {
            $(this).attr("class").split(" ").forEach((clazz) => {
                if (clazz.substr(0, "test-example-line-".length) === "test-example-line-") {
                    let end = clazz.substr("test-example-line-".length);
                    if (end !== "even" && end !== "odd" && end !== "0") {
                        showTestCaseMarker($(this).closest(".sample-test").find(".input"), "input", clazz, end);
                        showTestCaseMarker($(this).closest(".sample-test").find(".output"), "output", clazz, end);
                    }
                }
            });
        }, function() {
            $(this).attr("class").split(" ").forEach((clazz) => {
                if (clazz.substr(0, "test-example-line-".length) === "test-example-line-") {
                    let end = clazz.substr("test-example-line-".length);
                    if (end !== "even" && end !== "odd" && end !== "0") {
                        $("." + clazz).css("background-color", "");
                        $(this).closest(".problemindexholder").find(".testCaseMarker_input_" + end).hide();
                        $(this).closest(".problemindexholder").find(".testCaseMarker_output_" + end).hide();
                    }
                }
            });
        });

    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="B" data-uuid="ps_59469960d024700a759901135e75e10a3889e7d5">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">B. Print First 5 Alphabets</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>Your task is to print the first five letters of the English alphabet.</p></div><div class="input-specification"><div class="section-title">Input</div><p>There is no input for this problem.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print the first five uppercase English letters, one per line.</p></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id0016879060536658186" id="id0015476631280985043" class="input-output-copier">Copy</div></div><pre id="id0016879060536658186">There is no input.
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id0091402578967576" id="id006739864544172006" class="input-output-copier">Copy</div></div><pre id="id0091402578967576">A
B
C
D
E
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="C" data-uuid="ps_a635fcf21155973b98d6503d175e561a0b129207">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">C. Triangle</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>Your task is to print a right-angled triangle pattern using asterisks (<span class="tex-font-style-tt">*</span>).</p></div><div class="input-specification"><div class="section-title">Input</div><p>There is no input for this problem.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print the following pattern exactly as shown: </p><pre class="verbatim"><br>*****<br>****<br>***<br>**<br>*<br></pre></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id0029839156278767176" id="id001947521649532119" class="input-output-copier">Copy</div></div><pre id="id0029839156278767176">There is no input.
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id0047243163571177416" id="id008480975991255462" class="input-output-copier">Copy</div></div><pre id="id0047243163571177416">*****
****
***
**
*
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="D" data-uuid="ps_fabcbe06345237776c9cb4a29d7aa3955b162cc9">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">D. Print Z</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>Your task is to print the letter <span class="tex-font-style-tt">Z</span> using asterisks (<span class="tex-font-style-tt">*</span>) as shown below.</p></div><div class="input-specification"><div class="section-title">Input</div><p>There is no input for this problem.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print the following pattern exactly as shown: </p><pre class="verbatim"><br>*****<br>   *<br>  *<br> *<br>*****<br></pre></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id000010615138496328225" id="id008240130337395292" class="input-output-copier">Copy</div></div><pre id="id000010615138496328225">There is no input.
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id007082550365650344" id="id008664150217953653" class="input-output-copier">Copy</div></div><pre id="id007082550365650344">*****
   *
  *
 *
*****
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="E" data-uuid="ps_dc8c8f8ebc2bb23b5ad4aa9c56ec56ee7ef51fe9">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">E. Print Table of 5</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>Your task is to print the multiplication table of the <span class="tex-font-style-tt">5</span>.</p></div><div class="input-specification"><div class="section-title">Input</div><p>There is no input for this problem.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print the table of 5 in the format as shown in the example.</p></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id007495888064867978" id="id004118310775195757" class="input-output-copier">Copy</div></div><pre id="id007495888064867978">There is no input.
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id00938463840209552" id="id0016771979449201468" class="input-output-copier">Copy</div></div><pre id="id00938463840209552">5 * 1 = 5
5 * 2 = 10
5 * 3 = 15
5 * 4 = 20
5 * 5 = 25
5 * 6 = 30
5 * 7 = 35
5 * 8 = 40
5 * 9 = 45
5 * 10 = 50
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="F" data-uuid="ps_d17b3e81346ae050f6bf958baf13820b9001f976">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">F. Rectangle</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given the length and breadth of a rectangle. Your task is to calculate its area and perimeter.</p><p>The formulas are:</p><ul> <li> Area <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-1-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo&gt;=&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-1" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.642em, 1000.71em, 2.228em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-2"><span class="mo" id="MathJax-Span-3" style="font-family: MathJax_Main;">=</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: 0.075em; border-left: 0px solid; width: 0px; height: 0.432em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo>=</mo></math></span></span><script type="math/tex" id="MathJax-Element-1">=</script> length <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-2-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo&gt;&amp;#x00D7;&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-4" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.525em, 1000.59em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-5"><span class="mo" id="MathJax-Span-6" style="font-family: MathJax_Main;">×</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 0.718em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo>×</mo></math></span></span><script type="math/tex" id="MathJax-Element-2">\times</script> breadth </li><li> Perimeter <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-3-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo&gt;=&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-7" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.642em, 1000.71em, 2.228em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-8"><span class="mo" id="MathJax-Span-9" style="font-family: MathJax_Main;">=</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: 0.075em; border-left: 0px solid; width: 0px; height: 0.432em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo>=</mo></math></span></span><script type="math/tex" id="MathJax-Element-3">=</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-4-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;2&lt;/mn&gt;&lt;mo&gt;&amp;#x00D7;&lt;/mo&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mtext&gt;length&lt;/mtext&gt;&lt;mo&gt;+&lt;/mo&gt;&lt;mtext&gt;breadth&lt;/mtext&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-10" style="width: 12.064em; display: inline-block;"><span style="display: inline-block; position: relative; width: 9.898em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.232em, 1009.78em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-11"><span class="mn" id="MathJax-Span-12" style="font-family: MathJax_Main;">2</span><span class="mo" id="MathJax-Span-13" style="font-family: MathJax_Main; padding-left: 0.237em;">×</span><span class="mo" id="MathJax-Span-14" style="font-family: MathJax_Main; padding-left: 0.237em;">(</span><span class="mtext" id="MathJax-Span-15" style="font-family: MathJax_Main;">length</span><span class="mo" id="MathJax-Span-16" style="font-family: MathJax_Main; padding-left: 0.237em;">+</span><span class="mtext" id="MathJax-Span-17" style="font-family: MathJax_Main; padding-left: 0.237em;">breadth</span><span class="mo" id="MathJax-Span-18" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.361em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>2</mn><mo>×</mo><mo stretchy="false">(</mo><mtext>length</mtext><mo>+</mo><mtext>breadth</mtext><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-4">2 \times ( \text{length} + \text{breadth} )</script> </li></ul></div><div class="input-specification"><div class="section-title">Input</div><p>A single line containing two integers <span class="tex-font-style-tt">length</span> and <span class="tex-font-style-tt">breadth</span> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-5-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mrow class=&quot;MJX-TeXAtom-ORD&quot;&gt;&lt;mtext mathvariant=&quot;monospace&quot;&gt;length&lt;/mtext&gt;&lt;/mrow&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mrow class=&quot;MJX-TeXAtom-ORD&quot;&gt;&lt;mtext mathvariant=&quot;monospace&quot;&gt;breadth&lt;/mtext&gt;&lt;/mrow&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mn&gt;1000&lt;/mn&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-19" style="width: 16.221em; display: inline-block;"><span style="display: inline-block; position: relative; width: 13.293em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.232em, 1013.18em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-20"><span class="mo" id="MathJax-Span-21" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-22" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-23" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="texatom" id="MathJax-Span-24" style="padding-left: 0.296em;"><span class="mrow" id="MathJax-Span-25"><span class="mtext" id="MathJax-Span-26" style="font-family: MathJax_Typewriter;">length</span></span></span><span class="mo" id="MathJax-Span-27" style="font-family: MathJax_Main;">,</span><span class="texatom" id="MathJax-Span-28" style="padding-left: 0.179em;"><span class="mrow" id="MathJax-Span-29"><span class="mtext" id="MathJax-Span-30" style="font-family: MathJax_Typewriter;">breadth</span></span></span><span class="mo" id="MathJax-Span-31" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mn" id="MathJax-Span-32" style="font-family: MathJax_Main; padding-left: 0.296em;">1000</span><span class="mo" id="MathJax-Span-33" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.361em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mrow class="MJX-TeXAtom-ORD"><mtext mathvariant="monospace">length</mtext></mrow><mo>,</mo><mrow class="MJX-TeXAtom-ORD"><mtext mathvariant="monospace">breadth</mtext></mrow><mo>≤</mo><mn>1000</mn><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-5">(1 \le \texttt{length}, \texttt{breadth} \le 1000)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>You are given the length and breadth of a rectangle. Your task is to calculate its area and perimeter.</p><p>The formulas are:</p><ul> <li> Area <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-6-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo&gt;=&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-34" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.642em, 1000.71em, 2.228em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-35"><span class="mo" id="MathJax-Span-36" style="font-family: MathJax_Main;">=</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: 0.075em; border-left: 0px solid; width: 0px; height: 0.432em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo>=</mo></math></span></span><script type="math/tex" id="MathJax-Element-6">=</script> length <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-7-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo&gt;&amp;#x00D7;&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-37" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.525em, 1000.59em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-38"><span class="mo" id="MathJax-Span-39" style="font-family: MathJax_Main;">×</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 0.718em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo>×</mo></math></span></span><script type="math/tex" id="MathJax-Element-7">\times</script> breadth </li><li> Perimeter <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-8-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo&gt;=&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-40" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.642em, 1000.71em, 2.228em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-41"><span class="mo" id="MathJax-Span-42" style="font-family: MathJax_Main;">=</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: 0.075em; border-left: 0px solid; width: 0px; height: 0.432em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo>=</mo></math></span></span><script type="math/tex" id="MathJax-Element-8">=</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-9-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;2&lt;/mn&gt;&lt;mo&gt;&amp;#x00D7;&lt;/mo&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mtext&gt;length&lt;/mtext&gt;&lt;mo&gt;+&lt;/mo&gt;&lt;mtext&gt;breadth&lt;/mtext&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-43" style="width: 12.064em; display: inline-block;"><span style="display: inline-block; position: relative; width: 9.898em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.232em, 1009.78em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-44"><span class="mn" id="MathJax-Span-45" style="font-family: MathJax_Main;">2</span><span class="mo" id="MathJax-Span-46" style="font-family: MathJax_Main; padding-left: 0.237em;">×</span><span class="mo" id="MathJax-Span-47" style="font-family: MathJax_Main; padding-left: 0.237em;">(</span><span class="mtext" id="MathJax-Span-48" style="font-family: MathJax_Main;">length</span><span class="mo" id="MathJax-Span-49" style="font-family: MathJax_Main; padding-left: 0.237em;">+</span><span class="mtext" id="MathJax-Span-50" style="font-family: MathJax_Main; padding-left: 0.237em;">breadth</span><span class="mo" id="MathJax-Span-51" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.361em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>2</mn><mo>×</mo><mo stretchy="false">(</mo><mtext>length</mtext><mo>+</mo><mtext>breadth</mtext><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-9">2 \times ( \text{length} + \text{breadth} )</script> </li></ul></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id00553380807326886" id="id0048685872271057984" class="input-output-copier">Copy</div></div><pre id="id00553380807326886">5 7
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id009093068809887016" id="id00712814901848565" class="input-output-copier">Copy</div></div><pre id="id009093068809887016">Area = 35
Perimeter = 24
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="G" data-uuid="ps_a7eabebf9bead1120a182c8819053eb1bac010b7">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">G. Print Table of N</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given an integer <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-10-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-52" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-53"><span class="mi" id="MathJax-Span-54" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-10">N</script>. Your task is to print the multiplication table of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-11-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-55" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-56"><span class="mi" id="MathJax-Span-57" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-11">N</script> from <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-12-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-58" style="width: 0.647em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.53em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.35em, 1000.47em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-59"><span class="mn" id="MathJax-Span-60" style="font-family: MathJax_Main;">1</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 0.932em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>1</mn></math></span></span><script type="math/tex" id="MathJax-Element-12">1</script> to <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-13-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;10&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-61" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.35em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-62"><span class="mn" id="MathJax-Span-63" style="font-family: MathJax_Main;">10</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>10</mn></math></span></span><script type="math/tex" id="MathJax-Element-13">10</script>.</p></div><div class="input-specification"><div class="section-title">Input</div><p>A single integer <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-14-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-64" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-65"><span class="mi" id="MathJax-Span-66" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-14">N</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-15-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mn&gt;100&lt;/mn&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-67" style="width: 7.79em; display: inline-block;"><span style="display: inline-block; position: relative; width: 6.385em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.232em, 1006.27em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-68"><span class="mo" id="MathJax-Span-69" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-70" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-71" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-72" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-73" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mn" id="MathJax-Span-74" style="font-family: MathJax_Main; padding-left: 0.296em;">100</span><span class="mo" id="MathJax-Span-75" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.361em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mi>N</mi><mo>≤</mo><mn>100</mn><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-15">(1 \le N \le 100)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print the table of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-16-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-76" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-77"><span class="mi" id="MathJax-Span-78" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-16">N</script> in the format shown in the example below.</p></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id0013247100439696635" id="id00720215373195137" class="input-output-copier">Copy</div></div><pre id="id0013247100439696635">6
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id007781218661761473" id="id005576538443831885" class="input-output-copier">Copy</div></div><pre id="id007781218661761473">6 * 1 = 6
6 * 2 = 12
6 * 3 = 18
6 * 4 = 24
6 * 5 = 30
6 * 6 = 36
6 * 7 = 42
6 * 8 = 48
6 * 9 = 54
6 * 10 = 60
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="H1" data-uuid="ps_59d788b33a2f0690ab6af98150758ea10706d019">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">H1. Calculator</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-17-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-79" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-80"><span class="mi" id="MathJax-Span-81" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-17">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-18-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-82" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-83"><span class="mi" id="MathJax-Span-84" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-18">M</script>. Your task is to compute and print the results of the following operations:</p><ul> <li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-19-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;+&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-85" style="width: 3.867em; display: inline-block;"><span style="display: inline-block; position: relative; width: 3.165em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1003.16em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-86"><span class="mi" id="MathJax-Span-87" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-88" style="font-family: MathJax_Main; padding-left: 0.237em;">+</span><span class="mi" id="MathJax-Span-89" style="font-family: MathJax_Math-italic; padding-left: 0.237em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.075em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo>+</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-19">N + M</script> </li><li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-20-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;&amp;#x2212;&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-90" style="width: 3.867em; display: inline-block;"><span style="display: inline-block; position: relative; width: 3.165em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1003.16em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-91"><span class="mi" id="MathJax-Span-92" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-93" style="font-family: MathJax_Main; padding-left: 0.237em;">−</span><span class="mi" id="MathJax-Span-94" style="font-family: MathJax_Math-italic; padding-left: 0.237em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.075em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo>−</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-20">N - M</script> </li><li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-21-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;&amp;#x00D7;&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-95" style="width: 3.867em; display: inline-block;"><span style="display: inline-block; position: relative; width: 3.165em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1003.16em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-96"><span class="mi" id="MathJax-Span-97" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-98" style="font-family: MathJax_Main; padding-left: 0.237em;">×</span><span class="mi" id="MathJax-Span-99" style="font-family: MathJax_Math-italic; padding-left: 0.237em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo>×</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-21">N \times M</script> </li><li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-22-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;&amp;#x00F7;&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-100" style="width: 3.867em; display: inline-block;"><span style="display: inline-block; position: relative; width: 3.165em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1003.16em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-101"><span class="mi" id="MathJax-Span-102" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-103" style="font-family: MathJax_Main; padding-left: 0.237em;">÷</span><span class="mi" id="MathJax-Span-104" style="font-family: MathJax_Math-italic; padding-left: 0.237em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo>÷</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-22">N \div M</script> (integer division) </li><li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-23-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo lspace=&quot;thickmathspace&quot; rspace=&quot;thickmathspace&quot;&gt;mod&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-105" style="width: 5.389em; display: inline-block;"><span style="display: inline-block; position: relative; width: 4.394em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1004.39em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-106"><span class="mi" id="MathJax-Span-107" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-108" style="font-family: MathJax_Main; padding-left: 0.296em; padding-right: 0.296em;">mod</span><span class="mi" id="MathJax-Span-109" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo lspace="thickmathspace" rspace="thickmathspace">mod</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-23">N \bmod M</script> </li></ul></div><div class="input-specification"><div class="section-title">Input</div><p>A single line containing two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-24-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-110" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-111"><span class="mi" id="MathJax-Span-112" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-24">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-25-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-113" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-114"><span class="mi" id="MathJax-Span-115" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-25">M</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-26-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mn&gt;1000&lt;/mn&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-116" style="width: 10.249em; display: inline-block;"><span style="display: inline-block; position: relative; width: 8.375em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.232em, 1008.26em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-117"><span class="mo" id="MathJax-Span-118" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-119" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-120" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-121" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-122" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-123" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-124" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mn" id="MathJax-Span-125" style="font-family: MathJax_Main; padding-left: 0.296em;">1000</span><span class="mo" id="MathJax-Span-126" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.361em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mi>N</mi><mo>,</mo><mi>M</mi><mo>≤</mo><mn>1000</mn><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-26">(1 \le N, M \le 1000)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print five lines in the format as shown in the example.</p></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id009576120612123594" id="id00009252948688849627" class="input-output-copier">Copy</div></div><pre id="id009576120612123594">6 4
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id0003412854845471214" id="id0031157312083265587" class="input-output-copier">Copy</div></div><pre id="id0003412854845471214">6 + 4 = 10
6 - 4 = 2
6 * 4 = 24
6 / 4 = 1
6 % 4 = 2
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="H2" data-uuid="ps_e1574878f61f365ecb75490c3a19743e8c2d2d9e">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">H2. Calculator - II</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-27-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-127" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-128"><span class="mi" id="MathJax-Span-129" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-27">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-28-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-130" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-131"><span class="mi" id="MathJax-Span-132" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-28">M</script>. Your task is to compute and print the results of the following operations:</p><ul> <li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-29-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;+&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-133" style="width: 3.867em; display: inline-block;"><span style="display: inline-block; position: relative; width: 3.165em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1003.16em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-134"><span class="mi" id="MathJax-Span-135" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-136" style="font-family: MathJax_Main; padding-left: 0.237em;">+</span><span class="mi" id="MathJax-Span-137" style="font-family: MathJax_Math-italic; padding-left: 0.237em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.075em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo>+</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-29">N + M</script> </li><li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-30-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;&amp;#x2212;&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-138" style="width: 3.867em; display: inline-block;"><span style="display: inline-block; position: relative; width: 3.165em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1003.16em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-139"><span class="mi" id="MathJax-Span-140" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-141" style="font-family: MathJax_Main; padding-left: 0.237em;">−</span><span class="mi" id="MathJax-Span-142" style="font-family: MathJax_Math-italic; padding-left: 0.237em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.075em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo>−</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-30">N - M</script> </li><li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-31-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;&amp;#x00D7;&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-143" style="width: 3.867em; display: inline-block;"><span style="display: inline-block; position: relative; width: 3.165em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1003.16em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-144"><span class="mi" id="MathJax-Span-145" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-146" style="font-family: MathJax_Main; padding-left: 0.237em;">×</span><span class="mi" id="MathJax-Span-147" style="font-family: MathJax_Math-italic; padding-left: 0.237em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo>×</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-31">N \times M</script> </li><li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-32-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;&amp;#x00F7;&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-148" style="width: 3.867em; display: inline-block;"><span style="display: inline-block; position: relative; width: 3.165em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1003.16em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-149"><span class="mi" id="MathJax-Span-150" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-151" style="font-family: MathJax_Main; padding-left: 0.237em;">÷</span><span class="mi" id="MathJax-Span-152" style="font-family: MathJax_Math-italic; padding-left: 0.237em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo>÷</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-32">N \div M</script> (integer division) </li><li> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-33-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo lspace=&quot;thickmathspace&quot; rspace=&quot;thickmathspace&quot;&gt;mod&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-153" style="width: 5.389em; display: inline-block;"><span style="display: inline-block; position: relative; width: 4.394em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1004.39em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-154"><span class="mi" id="MathJax-Span-155" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-156" style="font-family: MathJax_Main; padding-left: 0.296em; padding-right: 0.296em;">mod</span><span class="mi" id="MathJax-Span-157" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi><mo lspace="thickmathspace" rspace="thickmathspace">mod</mo><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-33">N \bmod M</script> </li></ul></div><div class="input-specification"><div class="section-title">Input</div><p>Two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-34-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-158" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-159"><span class="mi" id="MathJax-Span-160" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-34">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-35-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-161" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-162"><span class="mi" id="MathJax-Span-163" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-35">M</script>, each given on a separate line <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-36-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;msup&gt;&lt;mn&gt;10&lt;/mn&gt;&lt;mn&gt;9&lt;/mn&gt;&lt;/msup&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-164" style="width: 9.546em; display: inline-block;"><span style="display: inline-block; position: relative; width: 7.79em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.115em, 1007.67em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-165"><span class="mo" id="MathJax-Span-166" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-167" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-168" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-169" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-170" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-171" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-172" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="msubsup" id="MathJax-Span-173" style="padding-left: 0.296em;"><span style="display: inline-block; position: relative; width: 1.408em; height: 0px;"><span style="position: absolute; clip: rect(3.165em, 1000.94em, 4.16em, -999.997em); top: -3.978em; left: 0em;"><span class="mn" id="MathJax-Span-174" style="font-family: MathJax_Main;">10</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span><span style="position: absolute; top: -4.388em; left: 0.998em;"><span class="mn" id="MathJax-Span-175" style="font-size: 70.7%; font-family: MathJax_Main;">9</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span></span></span><span class="mo" id="MathJax-Span-176" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.504em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mi>N</mi><mo>,</mo><mi>M</mi><mo>≤</mo><msup><mn>10</mn><mn>9</mn></msup><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-36">(1 \le N, M \le 10^9)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print five results in the following format as shown in the example.</p></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id009327008685777739" id="id004603394768833836" class="input-output-copier">Copy</div></div><pre id="id009327008685777739">1000000000
1000000000
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id005311988608019466" id="id0042812864693764074" class="input-output-copier">Copy</div></div><pre id="id005311988608019466">1000000000 + 1000000000 = 2000000000

1000000000 - 1000000000 = 0

1000000000 * 1000000000 = 1000000000000000000

1000000000 / 1000000000 = 1

1000000000 % 1000000000 = 0
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="I" data-uuid="ps_cb190e51f0b7c39a5265ac88f565162e10c11c8d">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">I. Add Last Digits</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-37-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-177" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-178"><span class="mi" id="MathJax-Span-179" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-37">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-38-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-180" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-181"><span class="mi" id="MathJax-Span-182" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-38">M</script>. Your task is to find the sum of the last digits of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-39-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-183" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-184"><span class="mi" id="MathJax-Span-185" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-39">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-40-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-186" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-187"><span class="mi" id="MathJax-Span-188" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-40">M</script>.</p></div><div class="input-specification"><div class="section-title">Input</div><p>A single line containing two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-41-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-189" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-190"><span class="mi" id="MathJax-Span-191" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-41">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-42-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-192" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-193"><span class="mi" id="MathJax-Span-194" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-42">M</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-43-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mn&gt;1000&lt;/mn&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-195" style="width: 10.249em; display: inline-block;"><span style="display: inline-block; position: relative; width: 8.375em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.232em, 1008.26em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-196"><span class="mo" id="MathJax-Span-197" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-198" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-199" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-200" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-201" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-202" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-203" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mn" id="MathJax-Span-204" style="font-family: MathJax_Main; padding-left: 0.296em;">1000</span><span class="mo" id="MathJax-Span-205" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.361em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mi>N</mi><mo>,</mo><mi>M</mi><mo>≤</mo><mn>1000</mn><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-43">(1 \le N, M \le 1000)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print a single integer — the sum of the last digits of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-44-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-206" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-207"><span class="mi" id="MathJax-Span-208" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-44">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-45-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-209" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-210"><span class="mi" id="MathJax-Span-211" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-45">M</script>.</p></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id007257286401201445" id="id004618907169442191" class="input-output-copier">Copy</div></div><pre id="id007257286401201445">169 125
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id007957216283536072" id="id006946517064405717" class="input-output-copier">Copy</div></div><pre id="id007957216283536072">14
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="J" data-uuid="ps_5ab12dffc8c74835741ba569eecb1392f97d47bc">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">J. Even or Odd</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given an integer <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-46-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-212" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-213"><span class="mi" id="MathJax-Span-214" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-46">N</script>. Your task is to determine whether the number is even or odd.</p></div><div class="input-specification"><div class="section-title">Input</div><p>A single integer <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-47-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-215" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-216"><span class="mi" id="MathJax-Span-217" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-47">N</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-48-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;msup&gt;&lt;mn&gt;10&lt;/mn&gt;&lt;mn&gt;9&lt;/mn&gt;&lt;/msup&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-218" style="width: 7.673em; display: inline-block;"><span style="display: inline-block; position: relative; width: 6.268em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.115em, 1006.15em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-219"><span class="mo" id="MathJax-Span-220" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-221" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-222" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-223" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-224" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="msubsup" id="MathJax-Span-225" style="padding-left: 0.296em;"><span style="display: inline-block; position: relative; width: 1.408em; height: 0px;"><span style="position: absolute; clip: rect(3.165em, 1000.94em, 4.16em, -999.997em); top: -3.978em; left: 0em;"><span class="mn" id="MathJax-Span-226" style="font-family: MathJax_Main;">10</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span><span style="position: absolute; top: -4.388em; left: 0.998em;"><span class="mn" id="MathJax-Span-227" style="font-size: 70.7%; font-family: MathJax_Main;">9</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span></span></span><span class="mo" id="MathJax-Span-228" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.504em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mi>N</mi><mo>≤</mo><msup><mn>10</mn><mn>9</mn></msup><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-48">(1 \le N \le 10^9)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><ul> <li> <span class="tex-font-style-tt">Even</span> if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-49-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-229" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-230"><span class="mi" id="MathJax-Span-231" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-49">N</script> is even </li><li> <span class="tex-font-style-tt">Odd</span> if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-50-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-232" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-233"><span class="mi" id="MathJax-Span-234" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-50">N</script> is odd </li></ul></div><div class="sample-tests"><div class="section-title">Examples</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id0069660033085514" id="id002288714989706312" class="input-output-copier">Copy</div></div><pre id="id0069660033085514">20
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id009551882949830869" id="id003018037777714355" class="input-output-copier">Copy</div></div><pre id="id009551882949830869">Even
</pre></div><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id001038741777229506" id="id0020585005273799872" class="input-output-copier">Copy</div></div><pre id="id001038741777229506">3
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id003550628725538355" id="id003618862960979047" class="input-output-copier">Copy</div></div><pre id="id003550628725538355">Odd
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="K" data-uuid="ps_d237c2d5ed3b0a335c27676b8e876fc212379ab6">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">K. Factor</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-51-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-235" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-236"><span class="mi" id="MathJax-Span-237" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-51">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-52-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;F&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-238" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-239"><span class="mi" id="MathJax-Span-240" style="font-family: MathJax_Math-italic;">F<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.12em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>F</mi></math></span></span><script type="math/tex" id="MathJax-Element-52">F</script>. Your task is to check whether <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-53-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;F&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-241" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-242"><span class="mi" id="MathJax-Span-243" style="font-family: MathJax_Math-italic;">F<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.12em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>F</mi></math></span></span><script type="math/tex" id="MathJax-Element-53">F</script> is a factor of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-54-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-244" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-245"><span class="mi" id="MathJax-Span-246" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-54">N</script>.</p><p>A number <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-55-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;F&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-247" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-248"><span class="mi" id="MathJax-Span-249" style="font-family: MathJax_Math-italic;">F<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.12em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>F</mi></math></span></span><script type="math/tex" id="MathJax-Element-55">F</script> is said to be a factor of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-56-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-250" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-251"><span class="mi" id="MathJax-Span-252" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-56">N</script> if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-57-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-253" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-254"><span class="mi" id="MathJax-Span-255" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-57">N</script> is divisible by <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-58-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;F&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-256" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-257"><span class="mi" id="MathJax-Span-258" style="font-family: MathJax_Math-italic;">F<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.12em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>F</mi></math></span></span><script type="math/tex" id="MathJax-Element-58">F</script>.</p></div><div class="input-specification"><div class="section-title">Input</div><p>A single line containing two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-59-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-259" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-260"><span class="mi" id="MathJax-Span-261" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-59">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-60-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;F&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-262" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-263"><span class="mi" id="MathJax-Span-264" style="font-family: MathJax_Math-italic;">F<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.12em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>F</mi></math></span></span><script type="math/tex" id="MathJax-Element-60">F</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-61-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;F&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;msup&gt;&lt;mn&gt;10&lt;/mn&gt;&lt;mn&gt;9&lt;/mn&gt;&lt;/msup&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-265" style="width: 9.136em; display: inline-block;"><span style="display: inline-block; position: relative; width: 7.497em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.115em, 1007.38em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-266"><span class="mo" id="MathJax-Span-267" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-268" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-269" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-270" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-271" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-272" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">F<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.12em;"></span></span><span class="mo" id="MathJax-Span-273" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="msubsup" id="MathJax-Span-274" style="padding-left: 0.296em;"><span style="display: inline-block; position: relative; width: 1.408em; height: 0px;"><span style="position: absolute; clip: rect(3.165em, 1000.94em, 4.16em, -999.997em); top: -3.978em; left: 0em;"><span class="mn" id="MathJax-Span-275" style="font-family: MathJax_Main;">10</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span><span style="position: absolute; top: -4.388em; left: 0.998em;"><span class="mn" id="MathJax-Span-276" style="font-size: 70.7%; font-family: MathJax_Main;">9</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span></span></span><span class="mo" id="MathJax-Span-277" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.504em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mi>N</mi><mo>,</mo><mi>F</mi><mo>≤</mo><msup><mn>10</mn><mn>9</mn></msup><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-61">(1 \le N, F \le 10^9)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print: </p><ul> <li> <span class="tex-font-style-tt">Yes</span> if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-62-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;F&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-278" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-279"><span class="mi" id="MathJax-Span-280" style="font-family: MathJax_Math-italic;">F<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.12em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>F</mi></math></span></span><script type="math/tex" id="MathJax-Element-62">F</script> is a factor of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-63-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-281" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-282"><span class="mi" id="MathJax-Span-283" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-63">N</script> </li><li> <span class="tex-font-style-tt">No</span> otherwise </li></ul></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id009874212718394453" id="id000077116052159184045" class="input-output-copier">Copy</div></div><pre id="id009874212718394453">36 6
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id009963779496668089" id="id005518153385770934" class="input-output-copier">Copy</div></div><pre id="id009963779496668089">Yes
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="L" data-uuid="ps_c6cbd1dab61e1e67fa3093b4119bb618dd5e0093">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">L. Multiple</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-64-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-284" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-285"><span class="mi" id="MathJax-Span-286" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-64">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-65-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-287" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-288"><span class="mi" id="MathJax-Span-289" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-65">M</script>. Your task is to check whether <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-66-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-290" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-291"><span class="mi" id="MathJax-Span-292" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-66">M</script> is a multiple of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-67-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-293" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-294"><span class="mi" id="MathJax-Span-295" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-67">N</script>.</p><p>A number <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-68-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-296" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-297"><span class="mi" id="MathJax-Span-298" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-68">M</script> is said to be a multiple of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-69-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-299" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-300"><span class="mi" id="MathJax-Span-301" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-69">N</script> if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-70-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-302" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-303"><span class="mi" id="MathJax-Span-304" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-70">M</script> is divisible by <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-71-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-305" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-306"><span class="mi" id="MathJax-Span-307" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-71">N</script>.</p></div><div class="input-specification"><div class="section-title">Input</div><p>A single line containing two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-72-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-308" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-309"><span class="mi" id="MathJax-Span-310" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-72">N</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-73-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-311" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-312"><span class="mi" id="MathJax-Span-313" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-73">M</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-74-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;msup&gt;&lt;mn&gt;10&lt;/mn&gt;&lt;mn&gt;9&lt;/mn&gt;&lt;/msup&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-314" style="width: 9.546em; display: inline-block;"><span style="display: inline-block; position: relative; width: 7.79em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.115em, 1007.67em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-315"><span class="mo" id="MathJax-Span-316" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-317" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-318" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-319" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-320" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-321" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-322" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="msubsup" id="MathJax-Span-323" style="padding-left: 0.296em;"><span style="display: inline-block; position: relative; width: 1.408em; height: 0px;"><span style="position: absolute; clip: rect(3.165em, 1000.94em, 4.16em, -999.997em); top: -3.978em; left: 0em;"><span class="mn" id="MathJax-Span-324" style="font-family: MathJax_Main;">10</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span><span style="position: absolute; top: -4.388em; left: 0.998em;"><span class="mn" id="MathJax-Span-325" style="font-size: 70.7%; font-family: MathJax_Main;">9</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span></span></span><span class="mo" id="MathJax-Span-326" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.504em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mi>N</mi><mo>,</mo><mi>M</mi><mo>≤</mo><msup><mn>10</mn><mn>9</mn></msup><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-74">(1 \le N, M \le 10^9)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print: </p><ul> <li> <span class="tex-font-style-tt">Yes</span> if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-75-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;M&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-327" style="width: 1.291em; display: inline-block;"><span style="display: inline-block; position: relative; width: 1.057em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1001.06em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-328"><span class="mi" id="MathJax-Span-329" style="font-family: MathJax_Math-italic;">M<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>M</mi></math></span></span><script type="math/tex" id="MathJax-Element-75">M</script> is a multiple of <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-76-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;N&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-330" style="width: 1.115em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.881em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.88em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-331"><span class="mi" id="MathJax-Span-332" style="font-family: MathJax_Math-italic;">N<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>N</mi></math></span></span><script type="math/tex" id="MathJax-Element-76">N</script> </li><li> <span class="tex-font-style-tt">No</span> otherwise </li></ul></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id009749520604581837" id="id001317824933394801" class="input-output-copier">Copy</div></div><pre id="id009749520604581837">6 36
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id006203547375193199" id="id006144544544200509" class="input-output-copier">Copy</div></div><pre id="id006203547375193199">Yes
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="M" data-uuid="ps_a16261b7cdfb6c5a2b74336a404fb4c97ee481c4">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">M. Pass or Fail</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given the marks obtained by a student. Your task is to check whether the student has passed or failed.</p><p>A student is considered to have passed if the marks obtained are at least <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-77-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;35&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-333" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.35em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-334"><span class="mn" id="MathJax-Span-335" style="font-family: MathJax_Main;">35</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>35</mn></math></span></span><script type="math/tex" id="MathJax-Element-77">35</script>.</p></div><div class="input-specification"><div class="section-title">Input</div><p>A single integer <span class="tex-font-style-tt">marks</span> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-78-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mrow class=&quot;MJX-TeXAtom-ORD&quot;&gt;&lt;mtext mathvariant=&quot;monospace&quot;&gt;marks&lt;/mtext&gt;&lt;/mrow&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mn&gt;100&lt;/mn&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-336" style="width: 9.956em; display: inline-block;"><span style="display: inline-block; position: relative; width: 8.141em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.232em, 1008.02em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-337"><span class="mo" id="MathJax-Span-338" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-339" style="font-family: MathJax_Main;">0</span><span class="mo" id="MathJax-Span-340" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="texatom" id="MathJax-Span-341" style="padding-left: 0.296em;"><span class="mrow" id="MathJax-Span-342"><span class="mtext" id="MathJax-Span-343" style="font-family: MathJax_Typewriter;">marks</span></span></span><span class="mo" id="MathJax-Span-344" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mn" id="MathJax-Span-345" style="font-family: MathJax_Main; padding-left: 0.296em;">100</span><span class="mo" id="MathJax-Span-346" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.361em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>0</mn><mo>≤</mo><mrow class="MJX-TeXAtom-ORD"><mtext mathvariant="monospace">marks</mtext></mrow><mo>≤</mo><mn>100</mn><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-78">(0 \le \texttt{marks} \le 100)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print: </p><ul> <li> <span class="tex-font-style-tt">Pass</span> if the student has passed </li><li> <span class="tex-font-style-tt">Fail</span> otherwise </li></ul></div><div class="sample-tests"><div class="section-title">Examples</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id00853095637097057" id="id0019310055130250237" class="input-output-copier">Copy</div></div><pre id="id00853095637097057">36
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id006252744593523739" id="id0020512017334108035" class="input-output-copier">Copy</div></div><pre id="id006252744593523739">Pass
</pre></div><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id004741747537946731" id="id007230871296698707" class="input-output-copier">Copy</div></div><pre id="id004741747537946731">20
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id008597178923116651" id="id00995888700810166" class="input-output-copier">Copy</div></div><pre id="id008597178923116651">Fail
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="N" data-uuid="ps_fd7289bc26a6af4310c779364967959b2e6a8da9">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">N. Max and Min of 2 Numbers</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-79-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;A&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-347" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-348"><span class="mi" id="MathJax-Span-349" style="font-family: MathJax_Math-italic;">A</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>A</mi></math></span></span><script type="math/tex" id="MathJax-Element-79">A</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-80-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;B&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-350" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-351"><span class="mi" id="MathJax-Span-352" style="font-family: MathJax_Math-italic;">B</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>B</mi></math></span></span><script type="math/tex" id="MathJax-Element-80">B</script>. Your task is to find the minimum and maximum among them.</p></div><div class="input-specification"><div class="section-title">Input</div><p>A single line containing two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-81-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;A&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-353" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-354"><span class="mi" id="MathJax-Span-355" style="font-family: MathJax_Math-italic;">A</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>A</mi></math></span></span><script type="math/tex" id="MathJax-Element-81">A</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-82-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;B&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-356" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-357"><span class="mi" id="MathJax-Span-358" style="font-family: MathJax_Math-italic;">B</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>B</mi></math></span></span><script type="math/tex" id="MathJax-Element-82">B</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-83-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;A&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;B&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;msup&gt;&lt;mn&gt;10&lt;/mn&gt;&lt;mn&gt;9&lt;/mn&gt;&lt;/msup&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-359" style="width: 9.019em; display: inline-block;"><span style="display: inline-block; position: relative; width: 7.38em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.115em, 1007.26em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-360"><span class="mo" id="MathJax-Span-361" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-362" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-363" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-364" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">A</span><span class="mo" id="MathJax-Span-365" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-366" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">B</span><span class="mo" id="MathJax-Span-367" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="msubsup" id="MathJax-Span-368" style="padding-left: 0.296em;"><span style="display: inline-block; position: relative; width: 1.408em; height: 0px;"><span style="position: absolute; clip: rect(3.165em, 1000.94em, 4.16em, -999.997em); top: -3.978em; left: 0em;"><span class="mn" id="MathJax-Span-369" style="font-family: MathJax_Main;">10</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span><span style="position: absolute; top: -4.388em; left: 0.998em;"><span class="mn" id="MathJax-Span-370" style="font-size: 70.7%; font-family: MathJax_Main;">9</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span></span></span><span class="mo" id="MathJax-Span-371" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.504em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mi>A</mi><mo>,</mo><mi>B</mi><mo>≤</mo><msup><mn>10</mn><mn>9</mn></msup><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-83">(1 \le A, B \le 10^9)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print two lines: </p><ul> <li> <span class="tex-font-style-tt">Min = X</span> </li><li> <span class="tex-font-style-tt">Max = Y</span> </li></ul><p>where <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-84-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-372" style="width: 0.998em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.823em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.82em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-373"><span class="mi" id="MathJax-Span-374" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi></math></span></span><script type="math/tex" id="MathJax-Element-84">X</script> is the minimum value and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-85-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-375" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-376"><span class="mi" id="MathJax-Span-377" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi></math></span></span><script type="math/tex" id="MathJax-Element-85">Y</script> is the maximum value.</p></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id002297860272062484" id="id0021304726341253433" class="input-output-copier">Copy</div></div><pre id="id002297860272062484">12 9
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id007846397910491641" id="id007573337780569194" class="input-output-copier">Copy</div></div><pre id="id007846397910491641">Min = 9
Max = 12
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="O" data-uuid="ps_d13dce01806abeee4a31b00aedef8b4d3d7ecffb">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">O. Max and Min of 3 Numbers</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given three integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-86-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;A&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-378" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-379"><span class="mi" id="MathJax-Span-380" style="font-family: MathJax_Math-italic;">A</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>A</mi></math></span></span><script type="math/tex" id="MathJax-Element-86">A</script>, <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-87-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;B&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-381" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-382"><span class="mi" id="MathJax-Span-383" style="font-family: MathJax_Math-italic;">B</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>B</mi></math></span></span><script type="math/tex" id="MathJax-Element-87">B</script>, and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-88-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;C&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-384" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-385"><span class="mi" id="MathJax-Span-386" style="font-family: MathJax_Math-italic;">C<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>C</mi></math></span></span><script type="math/tex" id="MathJax-Element-88">C</script>. Your task is to find the minimum and maximum among them.</p></div><div class="input-specification"><div class="section-title">Input</div><p>A single line containing three integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-89-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;A&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-387" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-388"><span class="mi" id="MathJax-Span-389" style="font-family: MathJax_Math-italic;">A</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>A</mi></math></span></span><script type="math/tex" id="MathJax-Element-89">A</script>, <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-90-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;B&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-390" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-391"><span class="mi" id="MathJax-Span-392" style="font-family: MathJax_Math-italic;">B</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>B</mi></math></span></span><script type="math/tex" id="MathJax-Element-90">B</script>, and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-91-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;C&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-393" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-394"><span class="mi" id="MathJax-Span-395" style="font-family: MathJax_Math-italic;">C<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>C</mi></math></span></span><script type="math/tex" id="MathJax-Element-91">C</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-92-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;1&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;A&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;B&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;C&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;msup&gt;&lt;mn&gt;10&lt;/mn&gt;&lt;mn&gt;9&lt;/mn&gt;&lt;/msup&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-396" style="width: 10.542em; display: inline-block;"><span style="display: inline-block; position: relative; width: 8.609em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.115em, 1008.49em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-397"><span class="mo" id="MathJax-Span-398" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-399" style="font-family: MathJax_Main;">1</span><span class="mo" id="MathJax-Span-400" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-401" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">A</span><span class="mo" id="MathJax-Span-402" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-403" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">B</span><span class="mo" id="MathJax-Span-404" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-405" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">C<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.061em;"></span></span><span class="mo" id="MathJax-Span-406" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="msubsup" id="MathJax-Span-407" style="padding-left: 0.296em;"><span style="display: inline-block; position: relative; width: 1.408em; height: 0px;"><span style="position: absolute; clip: rect(3.165em, 1000.94em, 4.16em, -999.997em); top: -3.978em; left: 0em;"><span class="mn" id="MathJax-Span-408" style="font-family: MathJax_Main;">10</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span><span style="position: absolute; top: -4.388em; left: 0.998em;"><span class="mn" id="MathJax-Span-409" style="font-size: 70.7%; font-family: MathJax_Main;">9</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span></span></span><span class="mo" id="MathJax-Span-410" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.504em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>1</mn><mo>≤</mo><mi>A</mi><mo>,</mo><mi>B</mi><mo>,</mo><mi>C</mi><mo>≤</mo><msup><mn>10</mn><mn>9</mn></msup><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-92">(1 \le A, B, C \le 10^9)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print two lines: </p><ul> <li> <span class="tex-font-style-tt">Min = X</span> </li><li> <span class="tex-font-style-tt">Max = Y</span> </li></ul><p>where <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-93-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-411" style="width: 0.998em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.823em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.82em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-412"><span class="mi" id="MathJax-Span-413" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi></math></span></span><script type="math/tex" id="MathJax-Element-93">X</script> is the minimum value and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-94-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-414" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-415"><span class="mi" id="MathJax-Span-416" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi></math></span></span><script type="math/tex" id="MathJax-Element-94">Y</script> is the maximum value.</p></div><div class="sample-tests"><div class="section-title">Example</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id0047594231502979367" id="id005170524434195585" class="input-output-copier">Copy</div></div><pre id="id0047594231502979367">12 9 15
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id002814076810389238" id="id009931838761050567" class="input-output-copier">Copy</div></div><pre id="id002814076810389238">Min = 9
Max = 15
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 2em;">


<div class="problemindexholder" problemindex="P" data-uuid="ps_8b71ae467e766f3450610b6d446e9d07aaefdef3">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">P. Student Performance Evaluation</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given the marks obtained by a student. Based on the marks, display an appropriate performance message according to the following rules:</p><ul> <li> If marks are greater than <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-95-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;90&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-417" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.35em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-418"><span class="mn" id="MathJax-Span-419" style="font-family: MathJax_Main;">90</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>90</mn></math></span></span><script type="math/tex" id="MathJax-Element-95">90</script>, print <span class="tex-font-style-tt">Excellent</span> </li><li> Else if marks are greater than <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-96-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;80&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-420" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.35em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-421"><span class="mn" id="MathJax-Span-422" style="font-family: MathJax_Main;">80</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>80</mn></math></span></span><script type="math/tex" id="MathJax-Element-96">80</script> and less than or equal to <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-97-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;90&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-423" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.35em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-424"><span class="mn" id="MathJax-Span-425" style="font-family: MathJax_Main;">90</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>90</mn></math></span></span><script type="math/tex" id="MathJax-Element-97">90</script>, print <span class="tex-font-style-tt">Good</span> </li><li> Else if marks are greater than <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-98-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;70&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-426" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-427"><span class="mn" id="MathJax-Span-428" style="font-family: MathJax_Main;">70</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>70</mn></math></span></span><script type="math/tex" id="MathJax-Element-98">70</script> and less than or equal to <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-99-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;80&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-429" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.35em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-430"><span class="mn" id="MathJax-Span-431" style="font-family: MathJax_Main;">80</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>80</mn></math></span></span><script type="math/tex" id="MathJax-Element-99">80</script>, print <span class="tex-font-style-tt">Fair</span> </li><li> Else if marks are greater than <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-100-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;60&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-432" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.35em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-433"><span class="mn" id="MathJax-Span-434" style="font-family: MathJax_Main;">60</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>60</mn></math></span></span><script type="math/tex" id="MathJax-Element-100">60</script> and less than or equal to <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-101-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;70&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-435" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-436"><span class="mn" id="MathJax-Span-437" style="font-family: MathJax_Main;">70</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>70</mn></math></span></span><script type="math/tex" id="MathJax-Element-101">70</script>, print <span class="tex-font-style-tt">Meets Expectations</span> </li><li> Else (marks less than or equal to <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-102-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mn&gt;60&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-438" style="width: 1.232em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.998em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.35em, 1000.94em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-439"><span class="mn" id="MathJax-Span-440" style="font-family: MathJax_Main;">60</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mn>60</mn></math></span></span><script type="math/tex" id="MathJax-Element-102">60</script>), print <span class="tex-font-style-tt">Below Par</span> </li></ul></div><div class="input-specification"><div class="section-title">Input</div><p>A single integer <span class="tex-font-style-tt">marks</span> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-103-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mrow class=&quot;MJX-TeXAtom-ORD&quot;&gt;&lt;mtext mathvariant=&quot;monospace&quot;&gt;marks&lt;/mtext&gt;&lt;/mrow&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mn&gt;100&lt;/mn&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-441" style="width: 9.956em; display: inline-block;"><span style="display: inline-block; position: relative; width: 8.141em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.232em, 1008.02em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-442"><span class="mo" id="MathJax-Span-443" style="font-family: MathJax_Main;">(</span><span class="mn" id="MathJax-Span-444" style="font-family: MathJax_Main;">0</span><span class="mo" id="MathJax-Span-445" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="texatom" id="MathJax-Span-446" style="padding-left: 0.296em;"><span class="mrow" id="MathJax-Span-447"><span class="mtext" id="MathJax-Span-448" style="font-family: MathJax_Typewriter;">marks</span></span></span><span class="mo" id="MathJax-Span-449" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mn" id="MathJax-Span-450" style="font-family: MathJax_Main; padding-left: 0.296em;">100</span><span class="mo" id="MathJax-Span-451" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.361em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mn>0</mn><mo>≤</mo><mrow class="MJX-TeXAtom-ORD"><mtext mathvariant="monospace">marks</mtext></mrow><mo>≤</mo><mn>100</mn><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-103">(0 \le \texttt{marks} \le 100)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print a single line containing the performance message corresponding to the given marks.</p></div><div class="sample-tests"><div class="section-title">Examples</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id003405151808235436" id="id00015806754451649696" class="input-output-copier">Copy</div></div><pre id="id003405151808235436">85
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id0005144806196742535" id="id003604748934008073" class="input-output-copier">Copy</div></div><pre id="id0005144806196742535">Good
</pre></div><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id0006138339633842904" id="id0044262930663729017" class="input-output-copier">Copy</div></div><pre id="id0006138339633842904">99
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id001963403302171669" id="id0029457363977032447" class="input-output-copier">Copy</div></div><pre id="id001963403302171669">Excellent
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
        <div style="margin-bottom: 1em;">


<div class="problemindexholder" problemindex="Q" data-uuid="ps_7842f3c43765c62d5299739f566bbdefc50bbdcf">
    <div style="display: none; margin:1em 0;text-align: center; position: relative;" class="alert alert-info diff-notifier">
        <div>The problem statement has recently been changed. <a class="view-changes" href="#">View the changes.</a></div>
        <span class="diff-notifier-close" style="position: absolute; top: 0.2em; right: 0.3em; cursor: pointer; font-size: 1.4em;">×</span>
    </div>
        <div class="ttypography"><div class="problem-statement"><div class="header"><div class="title">Q. Find the location point</div><div class="time-limit"><div class="property-title">time limit per test</div>1 second</div><div class="memory-limit"><div class="property-title">memory limit per test</div>256 megabytes</div><div class="input-file input-standard"><div class="property-title">input</div>standard input</div><div class="output-file output-standard"><div class="property-title">output</div>standard output</div></div><div><p>You are given the coordinates of a point <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-104-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-452" style="width: 3.457em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.813em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.232em, 1002.7em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-453"><span class="mo" id="MathJax-Span-454" style="font-family: MathJax_Main;">(</span><span class="mi" id="MathJax-Span-455" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span><span class="mo" id="MathJax-Span-456" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-457" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span><span class="mo" id="MathJax-Span-458" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.361em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mi>X</mi><mo>,</mo><mi>Y</mi><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-104">(X, Y)</script> on a Cartesian plane. Your task is to determine the location of the point.</p><p>The possible locations are: </p><ul> <li> <span class="tex-font-style-tt">Origin</span> — if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-105-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;mo&gt;=&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-459" style="width: 3.282em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.696em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.64em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-460"><span class="mi" id="MathJax-Span-461" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span><span class="mo" id="MathJax-Span-462" style="font-family: MathJax_Main; padding-left: 0.296em;">=</span><span class="mn" id="MathJax-Span-463" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi><mo>=</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-105">X = 0</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-106-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;mo&gt;=&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-464" style="width: 3.223em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.638em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.58em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-465"><span class="mi" id="MathJax-Span-466" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span><span class="mo" id="MathJax-Span-467" style="font-family: MathJax_Main; padding-left: 0.296em;">=</span><span class="mn" id="MathJax-Span-468" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi><mo>=</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-106">Y = 0</script> </li><li> <span class="tex-font-style-tt">X axis</span> — if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-107-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;mo&gt;=&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-469" style="width: 3.223em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.638em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.58em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-470"><span class="mi" id="MathJax-Span-471" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span><span class="mo" id="MathJax-Span-472" style="font-family: MathJax_Main; padding-left: 0.296em;">=</span><span class="mn" id="MathJax-Span-473" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi><mo>=</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-107">Y = 0</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-108-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;mo&gt;&amp;#x2260;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-474" style="width: 3.282em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.696em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.64em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-475"><span class="mi" id="MathJax-Span-476" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span><span class="mo" id="MathJax-Span-477" style="font-family: MathJax_Main; padding-left: 0.296em;">≠</span><span class="mn" id="MathJax-Span-478" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.289em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi><mo>≠</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-108">X \ne 0</script> </li><li> <span class="tex-font-style-tt">Y axis</span> — if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-109-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;mo&gt;=&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-479" style="width: 3.282em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.696em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.64em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-480"><span class="mi" id="MathJax-Span-481" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span><span class="mo" id="MathJax-Span-482" style="font-family: MathJax_Main; padding-left: 0.296em;">=</span><span class="mn" id="MathJax-Span-483" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi><mo>=</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-109">X = 0</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-110-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;mo&gt;&amp;#x2260;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-484" style="width: 3.223em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.638em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.58em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-485"><span class="mi" id="MathJax-Span-486" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span><span class="mo" id="MathJax-Span-487" style="font-family: MathJax_Main; padding-left: 0.296em;">≠</span><span class="mn" id="MathJax-Span-488" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.289em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi><mo>≠</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-110">Y \ne 0</script> </li><li> <span class="tex-font-style-tt">1st Quadrant</span> — if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-111-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;mo&gt;&amp;gt;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-489" style="width: 3.282em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.696em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.64em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-490"><span class="mi" id="MathJax-Span-491" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span><span class="mo" id="MathJax-Span-492" style="font-family: MathJax_Main; padding-left: 0.296em;">&gt;</span><span class="mn" id="MathJax-Span-493" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi><mo>&gt;</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-111">X > 0</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-112-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;mo&gt;&amp;gt;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-494" style="width: 3.223em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.638em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.58em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-495"><span class="mi" id="MathJax-Span-496" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span><span class="mo" id="MathJax-Span-497" style="font-family: MathJax_Main; padding-left: 0.296em;">&gt;</span><span class="mn" id="MathJax-Span-498" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi><mo>&gt;</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-112">Y > 0</script> </li><li> <span class="tex-font-style-tt">2nd Quadrant</span> — if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-113-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;mo&gt;&amp;lt;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-499" style="width: 3.282em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.696em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.64em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-500"><span class="mi" id="MathJax-Span-501" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span><span class="mo" id="MathJax-Span-502" style="font-family: MathJax_Main; padding-left: 0.296em;">&lt;</span><span class="mn" id="MathJax-Span-503" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi><mo>&lt;</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-113">X < 0</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-114-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;mo&gt;&amp;gt;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-504" style="width: 3.223em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.638em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.58em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-505"><span class="mi" id="MathJax-Span-506" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span><span class="mo" id="MathJax-Span-507" style="font-family: MathJax_Main; padding-left: 0.296em;">&gt;</span><span class="mn" id="MathJax-Span-508" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi><mo>&gt;</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-114">Y > 0</script> </li><li> <span class="tex-font-style-tt">3rd Quadrant</span> — if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-115-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;mo&gt;&amp;lt;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-509" style="width: 3.282em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.696em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.64em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-510"><span class="mi" id="MathJax-Span-511" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span><span class="mo" id="MathJax-Span-512" style="font-family: MathJax_Main; padding-left: 0.296em;">&lt;</span><span class="mn" id="MathJax-Span-513" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi><mo>&lt;</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-115">X < 0</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-116-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;mo&gt;&amp;lt;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-514" style="width: 3.223em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.638em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.58em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-515"><span class="mi" id="MathJax-Span-516" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span><span class="mo" id="MathJax-Span-517" style="font-family: MathJax_Main; padding-left: 0.296em;">&lt;</span><span class="mn" id="MathJax-Span-518" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi><mo>&lt;</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-116">Y < 0</script> </li><li> <span class="tex-font-style-tt">4th Quadrant</span> — if <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-117-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;mo&gt;&amp;gt;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-519" style="width: 3.282em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.696em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.64em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-520"><span class="mi" id="MathJax-Span-521" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span><span class="mo" id="MathJax-Span-522" style="font-family: MathJax_Main; padding-left: 0.296em;">&gt;</span><span class="mn" id="MathJax-Span-523" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi><mo>&gt;</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-117">X > 0</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-118-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;mo&gt;&amp;lt;&lt;/mo&gt;&lt;mn&gt;0&lt;/mn&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-524" style="width: 3.223em; display: inline-block;"><span style="display: inline-block; position: relative; width: 2.638em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1002.58em, 2.403em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-525"><span class="mi" id="MathJax-Span-526" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span><span class="mo" id="MathJax-Span-527" style="font-family: MathJax_Main; padding-left: 0.296em;">&lt;</span><span class="mn" id="MathJax-Span-528" style="font-family: MathJax_Main; padding-left: 0.296em;">0</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.139em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi><mo>&lt;</mo><mn>0</mn></math></span></span><script type="math/tex" id="MathJax-Element-118">Y < 0</script> </li></ul></div><div class="input-specification"><div class="section-title">Input</div><p>A single line containing two integers <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-119-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-529" style="width: 0.998em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.823em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.82em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-530"><span class="mi" id="MathJax-Span-531" style="font-family: MathJax_Math-italic;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>X</mi></math></span></span><script type="math/tex" id="MathJax-Element-119">X</script> and <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-120-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-532" style="width: 0.94em; display: inline-block;"><span style="display: inline-block; position: relative; width: 0.764em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.291em, 1000.76em, 2.345em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-533"><span class="mi" id="MathJax-Span-534" style="font-family: MathJax_Math-italic;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.068em; border-left: 0px solid; width: 0px; height: 1.004em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mi>Y</mi></math></span></span><script type="math/tex" id="MathJax-Element-120">Y</script> <span class="MathJax_Preview" style="color: inherit;"></span><span class="MathJax" id="MathJax-Element-121-Frame" tabindex="0" style="position: relative;" data-mathml="&lt;math xmlns=&quot;http://www.w3.org/1998/Math/MathML&quot;&gt;&lt;mo stretchy=&quot;false&quot;&gt;(&lt;/mo&gt;&lt;mo&gt;&amp;#x2212;&lt;/mo&gt;&lt;msup&gt;&lt;mn&gt;10&lt;/mn&gt;&lt;mn&gt;9&lt;/mn&gt;&lt;/msup&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;mi&gt;X&lt;/mi&gt;&lt;mo&gt;,&lt;/mo&gt;&lt;mi&gt;Y&lt;/mi&gt;&lt;mo&gt;&amp;#x2264;&lt;/mo&gt;&lt;msup&gt;&lt;mn&gt;10&lt;/mn&gt;&lt;mn&gt;9&lt;/mn&gt;&lt;/msup&gt;&lt;mo stretchy=&quot;false&quot;&gt;)&lt;/mo&gt;&lt;/math&gt;" role="presentation"><nobr aria-hidden="true"><span class="math" id="MathJax-Span-535" style="width: 11.186em; display: inline-block;"><span style="display: inline-block; position: relative; width: 9.136em; height: 0px; font-size: 122%;"><span style="position: absolute; clip: rect(1.115em, 1009.02em, 2.579em, -999.997em); top: -2.163em; left: 0em;"><span class="mrow" id="MathJax-Span-536"><span class="mo" id="MathJax-Span-537" style="font-family: MathJax_Main;">(</span><span class="mo" id="MathJax-Span-538" style="font-family: MathJax_Main;">−</span><span class="msubsup" id="MathJax-Span-539"><span style="display: inline-block; position: relative; width: 1.408em; height: 0px;"><span style="position: absolute; clip: rect(3.165em, 1000.94em, 4.16em, -999.997em); top: -3.978em; left: 0em;"><span class="mn" id="MathJax-Span-540" style="font-family: MathJax_Main;">10</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span><span style="position: absolute; top: -4.388em; left: 0.998em;"><span class="mn" id="MathJax-Span-541" style="font-size: 70.7%; font-family: MathJax_Main;">9</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span></span></span><span class="mo" id="MathJax-Span-542" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="mi" id="MathJax-Span-543" style="font-family: MathJax_Math-italic; padding-left: 0.296em;">X<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.003em;"></span></span><span class="mo" id="MathJax-Span-544" style="font-family: MathJax_Main;">,</span><span class="mi" id="MathJax-Span-545" style="font-family: MathJax_Math-italic; padding-left: 0.179em;">Y<span style="display: inline-block; overflow: hidden; height: 1px; width: 0.179em;"></span></span><span class="mo" id="MathJax-Span-546" style="font-family: MathJax_Main; padding-left: 0.296em;">≤</span><span class="msubsup" id="MathJax-Span-547" style="padding-left: 0.296em;"><span style="display: inline-block; position: relative; width: 1.408em; height: 0px;"><span style="position: absolute; clip: rect(3.165em, 1000.94em, 4.16em, -999.997em); top: -3.978em; left: 0em;"><span class="mn" id="MathJax-Span-548" style="font-family: MathJax_Main;">10</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span><span style="position: absolute; top: -4.388em; left: 0.998em;"><span class="mn" id="MathJax-Span-549" style="font-size: 70.7%; font-family: MathJax_Main;">9</span><span style="display: inline-block; width: 0px; height: 3.984em;"></span></span></span></span><span class="mo" id="MathJax-Span-550" style="font-family: MathJax_Main;">)</span></span><span style="display: inline-block; width: 0px; height: 2.169em;"></span></span></span><span style="display: inline-block; overflow: hidden; vertical-align: -0.354em; border-left: 0px solid; width: 0px; height: 1.504em;"></span></span></nobr><span class="MJX_Assistive_MathML" role="presentation"><math xmlns="http://www.w3.org/1998/Math/MathML"><mo stretchy="false">(</mo><mo>−</mo><msup><mn>10</mn><mn>9</mn></msup><mo>≤</mo><mi>X</mi><mo>,</mo><mi>Y</mi><mo>≤</mo><msup><mn>10</mn><mn>9</mn></msup><mo stretchy="false">)</mo></math></span></span><script type="math/tex" id="MathJax-Element-121">(-10^9 \le X, Y \le 10^9)</script>.</p></div><div class="output-specification"><div class="section-title">Output</div><p>Print a single line indicating the location of the point.</p></div><div class="sample-tests"><div class="section-title">Examples</div><div class="sample-test"><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id008190911005552309" id="id005850244241349627" class="input-output-copier">Copy</div></div><pre id="id008190911005552309">1 10
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id001629920558377307" id="id00666446960920405" class="input-output-copier">Copy</div></div><pre id="id001629920558377307">1st Quadrant
</pre></div><div class="input"><div class="title">Input<div title="Copy" data-clipboard-target="#id0009656439345106471" id="id004058347888219115" class="input-output-copier">Copy</div></div><pre id="id0009656439345106471">0 0
</pre></div><div class="output"><div class="title">Output<div title="Copy" data-clipboard-target="#id003702725770885955" id="id00520926287612364" class="input-output-copier">Copy</div></div><pre id="id003702725770885955">Origin
</pre></div></div></div></div><p>  </p></div>
</div>



<script type="text/javascript">
    $(document).ready(function () {

        const inputFileDiv = $("div.input-file");
        const inputFile = inputFileDiv.text();
        const outputFileDiv = $("div.output-file");
        const outputFile = outputFileDiv.text();


        if (!endsWith(inputFile, "standard input")
            && !endsWith(inputFile, "standard input")) {
            inputFileDiv.attr("style", "font-weight: bold");
        }

        if (!endsWith(outputFile, "standard output")
            && !endsWith(outputFile, "standard output")) {
            outputFileDiv.attr("style", "font-weight: bold");
        }

        const titleDiv = $("div.header div.title");




    });
</script>
        </div>
    </div>

    <div id="footer">
        <div><a href="https://codeforces.com/">Codeforces</a> (c) Copyright 2010-2026 Mike Mirzayanov</div>
        <div>The only programming contests Web 2.0 platform</div>

    </div>

</div>
</div>
    <script type="application/javascript">
        if ('serviceWorker' in navigator && 'fetch' in window && 'caches' in window) {
            navigator.serviceWorker.register('/service-worker-19965.js')
                .then(function (registration) {
                    console.log('Service worker registered: ', registration);
                })
                .catch(function (error) {
                    console.log('Registration failed: ', error);
                });
        }
    </script>
<script defer="" src="https://static.cloudflareinsights.com/beacon.min.js/vcd15cbe7772f49c399c6a5babf22c1241717689176015" integrity="sha512-ZpsOmlRQV6y907TI0dKBHq9Md29nnaEIPlkf84rnaERnq6zvWvPUqr2ft8M1aS28oN72PdrCzSjY4U6VaAw1EQ==" data-cf-beacon="{&quot;version&quot;:&quot;2024.11.0&quot;,&quot;token&quot;:&quot;316c29dfa3804effb5a91c52a59b5bd0&quot;,&quot;server_timing&quot;:{&quot;name&quot;:{&quot;cfCacheStatus&quot;:true,&quot;cfEdge&quot;:true,&quot;cfExtPri&quot;:true,&quot;cfL4&quot;:true,&quot;cfOrigin&quot;:true,&quot;cfSpeedBrain&quot;:true},&quot;location_startswith&quot;:null}}" crossorigin="anonymous"></script>


<div id="datepick-div" style="display: none;"></div><div id="cboxOverlay" style="display: none;"></div><div id="colorbox" class="" style="display: none; padding-bottom: 42px; padding-right: 42px;"><div id="cboxWrapper"><div><div id="cboxTopLeft" style="float: left;"></div><div id="cboxTopCenter" style="float: left;"></div><div id="cboxTopRight" style="float: left;"></div></div><div style="clear: left;"><div id="cboxMiddleLeft" style="float: left;"></div><div id="cboxContent" style="float: left;"><div id="cboxLoadedContent" style="width: 0px; height: 0px; overflow: hidden; float: left;"></div><div id="cboxLoadingOverlay" style="float: left;"></div><div id="cboxLoadingGraphic" style="float: left;"></div><div id="cboxTitle" style="float: left;"></div><div id="cboxCurrent" style="float: left;"></div><div id="cboxNext" style="float: left;"></div><div id="cboxPrevious" style="float: left;"></div><div id="cboxSlideshow" style="float: left;"></div><div id="cboxClose" style="float: left;"></div></div><div id="cboxMiddleRight" style="float: left;"></div></div><div style="clear: left;"><div id="cboxBottomLeft" style="float: left;"></div><div id="cboxBottomCenter" style="float: left;"></div><div id="cboxBottomRight" style="float: left;"></div></div></div><div style="position: absolute; width: 9999px; visibility: hidden; display: none;"></div></div><div style="position: absolute; width: 0px; height: 0px; overflow: hidden; padding: 0px; border: 0px; margin: 0px;"><div id="MathJax_Font_Test" style="position: absolute; visibility: hidden; top: 0px; left: 0px; width: auto; min-width: 0px; max-width: none; padding: 0px; border: 0px; margin: 0px; white-space: nowrap; text-align: left; text-indent: 0px; text-transform: none; line-height: normal; letter-spacing: normal; word-spacing: normal; font-size: 40px; font-weight: normal; font-style: normal; font-size-adjust: none; font-family: MathJax_Math-italic, sans-serif;"></div></div></body>
