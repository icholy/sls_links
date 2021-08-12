# sls_links

Find the lambda & cloudwatch links for a serverless service or provider.

## Install

``` sh
$ go install github.com/icholy/sls_links@latest
```

## Example:

```
$ sls_links -env=staging
Lambda:  compassdigital-order-staging-root
Link:    https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-root
LogLink: https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-root

Lambda:  compassdigital-order-staging-order
Link:    https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-order
LogLink: https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-order

Lambda:  compassdigital-order-staging-brandorders
Link:    https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-brandorders
LogLink: https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-brandorders

Lambda:  compassdigital-order-staging-customerorders
Link:    https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-customerorders
LogLink: https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-customerorders

Lambda:  compassdigital-order-staging-search_index
Link:    https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-search_index
LogLink: https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-search$252Findex
```