/*
 * Cryptexly - Copyleft of Xavier D. Johnson.
 * me at xavierdjohnson dot com
 * http://www.xavierdjohnson.net/
 *
 * See LICENSE.
 */
window.onerror = function (message, file, line, col, error) {
    var s = "Cryptexly version: " + VERSION + "\n" +
            "Browser: " + navigator.userAgent + "\n" +
            "Unhandled error on " + file + ":" + line + ":" + col + "\n" +
            "\n" + message + "\n\n" + 
            "---------------------------------------------------------------" + 
            "------------------------------------------------------------------------\n" + 
            "Please copy this message and report the bug here https://github.com/detroitcybersec/cryptexly/issues/";
    alert(s);
    return false;
};


