#include "httplib.h"

int main() {
    httplib::Server svr;

    svr.Get("/hi", [](const httplib::Request &req, httplib::Response &res) {
        res.set_content("Hello World!", "text/plain");
    });

    svr.listen("100.0.0.30", 8080);
    return 0;
}
