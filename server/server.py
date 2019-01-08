import time
import datetime as dt
import argparse
import grpc
from grtest import grtest_pb2
from grtest import grtest_pb2_grpc
from concurrent import futures


_ONE_DAY_IN_SECONDS = 60 * 60 * 24
parser = argparse.ArgumentParser(description='gRPC test.')
parser.add_argument('--listen', default='[::]:12345')

args = parser.parse_args()


class TestService(grtest_pb2_grpc.TestServiceServicer):

    def PingRequest(self, request, context):
        print("sends:", "pong=reply", f"aux={request.aux}", sep="\n\t")
        return grtest_pb2.Pong(pong="reply", aux=request.aux)

    def Stream(self, request, context):
        counter = 0
        prev_time = dt.datetime.now()
        while True:
            curr_time = dt.datetime.now()
            if curr_time > prev_time + dt.timedelta(seconds=3):
                print(f"{curr_time} stream sends:",
                      "pong=pong", f"aux={counter}", sep="\n\t")
                prev_time = curr_time
                yield grtest_pb2.Pong(pong="pong", aux=f"{counter}")
            counter += 1


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=1))
    grtest_pb2_grpc.add_TestServiceServicer_to_server(TestService(), server)
    server.add_insecure_port(args.listen)
    server.start()

    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
