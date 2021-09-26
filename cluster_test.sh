#!/bin/sh

# Creates the CAP, and tries creating pods with a) no palindrome in label keys
# and b) a palindrome in the label keys. The latter should fail.

CAP_RESOURCE_NAME='suse-kubewarden-test'
kubectl apply -f - << EOF
apiVersion: policies.kubewarden.io/v1alpha2
kind: ClusterAdmissionPolicy
metadata:
  name: ${CAP_RESOURCE_NAME}
spec:
  module: registry://ghcr.io/adamkpickering/suse-kubewarden:latest
  rules:
  - apiGroups: [""]
    apiVersions: ["v1"]
    resources: ["pods"]
    operations:
    - CREATE
  mutating: false
  settings: {}
EOF

# should succeed
NO_PALINDROME_POD_RESOURCE_NAME='no-palindrome'
kubectl apply -f - << EOF
---
apiVersion: v1
kind: Pod
metadata:
  name: ${NO_PALINDROME_POD_RESOURCE_NAME}
  labels:
    notapalindrome: testvalue
spec:
  containers:
    - command:
        - sleep
        - "3600"
      image: busybox
      name: pods-simple-container
EOF

# should fail
HAS_PALINDROME_POD_RESOURCE_NAME='has-palindrome'
kubectl apply -f - << EOF
---
apiVersion: v1
kind: Pod
metadata:
  name: ${HAS_PALINDROME_POD_RESOURCE_NAME}
  labels:
    radar: testvalue
spec:
  containers:
    - command:
        - sleep
        - "3600"
      image: busybox
      name: pods-simple-container
EOF
