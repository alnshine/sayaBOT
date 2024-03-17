import json

from typing import List

from http.server import HTTPServer
from http.server import BaseHTTPRequestHandler

from sayaBot import *

from attrs import asdict
from attrs import define

HOST = "127.0.0.1"
PORT = 8027

@define
class Entry:

    name: str
    some_list: List[int]

class Handler(BaseHTTPRequestHandler):
    def do_POST(self):
        
        # read incoming sent data
        data = self.rfile.read(self._sent_data_size)
        
        # do something with it ...
        response = self._process(data.decode("utf-8"))

        # perpare the (json) response
        jsonbytes = self._prepare_json_response(response)
                    
        # send the (json) response back ...
        self.wfile.write(jsonbytes)

    def _process(self, data: str) -> List[Entry]:
        return [
            Entry("employee"+str(i), [d["id"], d["salary"]])
            for i, d in enumerate(json.loads(data))
        ]

    def _prepare_json_response(self, response: List[Entry]) -> bytes:
        self.send_response(200)
        self.send_header("Content-type", "application/json")
        self.end_headers()
        jsonstr = json.dumps(
            response,
            indent=4,
            default=asdict
        )
        return jsonstr.encode()

    @property
    def _sent_data_size(self) -> int:
        return int(self.headers.get("Content-Length"))

server = HTTPServer((HOST, PORT), Handler)
server.serve_forever()
server.serve_close()