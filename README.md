# KubeFigure

Render configurations and secrets directly into your YAML's during runtime. 

k apply -> mutatingWebHook -> fetch data source -> render -> deploy 

# Supported Data Sources:

* Terraform (S3)

* Vault 

* Consul


# Usage example:

In terminal run with the file `example.yaml`: 

```bash
$go run main.go input --input=example.yaml
```

This config pulls data from vault, consul and terraform:

```yaml

sources: 
  - sourceType: vault 
    vault:
      address: <path-to-vault>
      authType: approle 
      auth: 
        approle: <approle>
      value:
        path: <path-to-secret-key>
        # optional: jsonpath to extract a key if the secret is a json file 
        valPath: $.token  
  - sourceType: terraform
    terraform:
      storageType: s3 
      storage:
        bucket: <bucket-name>
        region: <region>
      value: 
        key: <key to state file in bucket>
        # jsonpath to extract some output from the remote state, same as remote_state stanza in terraform 
        stateValuePath: $.outputs.ec2_instance_security_group_id.value
  - sourceType: consulkv
    consul:
      authType: http 
      auth:
        address: <consul-address>
        port: <consul-port>
      value:
        key: <kv path in consul>
        # optional: jsonpath to extract a key if the kv value is a json file 
        valPath: $.project_id
        # optional to filter kv 
        options: 
          datacenter: <consul-data-center-or-region>
```