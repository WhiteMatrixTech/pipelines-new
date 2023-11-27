#!/bin/bash

# This image should be in sync with Dockerfile.visualization.
IMAGE="dave917/tensorflow:latest"
# tensorflow/tfx default entrypoint is Apache BEAM, because Apache BEAM doesn't
# support custom entrypoint for now. We need to override with --entrypoint ""
# for other `docker run` usecase.
# https://github.com/tensorflow/tfx/blob/master/tfx/tools/docker/Dockerfile#L71
../../../../hack/update-requirements.sh $IMAGE <requirements.in >requirements.txt
