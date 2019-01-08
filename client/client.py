import argparse

import grpc
from grtest import grtest_pb2
from grtest import grtest_pb2_grpc
import datetime as dt

parser = argparse.ArgumentParser(description='gRPC test.')
parser.add_argument('--remote', default='localhost:12345')
parser.add_argument('--ping', default='ping data')
parser.add_argument('--aux', default='ping aux')

args = parser.parse_args()

channel = grpc.insecure_channel(args.remote)
stub = grtest_pb2_grpc.TestServiceStub(channel)

req = grtest_pb2.Ping(ping=args.ping, aux=args.aux)
reply = stub.PingRequest(req)
print("client sends:", end="\n\t")
print(f"ping={args.ping}", f"aux={args.aux}", sep="\n\t")
print("server receives:", end="\n\t")
print(f"pong={reply.pong}", f"aux={reply.aux}", sep="\n\t")


prev_time = dt.datetime.now()
for r in stub.Stream(req):
    curr_time = dt.datetime.now()
    if curr_time > prev_time + dt.timedelta(seconds=3):
        print(f"stream reply: {r.pong}:{r.aux}")
        prev_time = curr_time
