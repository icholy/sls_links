# sls_links

Find the lambda & cloudwatch links for a serverless service or provider.

## Install

``` sh
$ go install github.com/icholy/sls_links@latest
```

## Example:

```
$ sls_links ./compassdigital.service.order ./compassdigital.provider.order.cdl 
Name:   compassdigital-order-staging-search_index
Lambda: https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-search_index
Logs:   https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-search_index

Name:   compassdigital-order-staging-root
Lambda: https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-root
Logs:   https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-root

Name:   compassdigital-order-staging-order
Lambda: https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-order
Logs:   https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-order

Name:   compassdigital-order-staging-brandorders
Lambda: https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-brandorders
Logs:   https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-brandorders

Name:   compassdigital-order-staging-customerorders
Lambda: https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-order-staging-customerorders
Logs:   https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-order-staging-customerorders

Name:   compassdigital-provider-order-cdl-staging-dexai
Lambda: https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-provider-order-cdl-staging-dexai
Logs:   https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-provider-order-cdl-staging-dexai

Name:   compassdigital-provider-order-cdl-staging-catchall
Lambda: https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-provider-order-cdl-staging-catchall
Logs:   https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-provider-order-cdl-staging-catchall

Name:   compassdigital-provider-order-cdl-staging-neworder
Lambda: https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/compassdigital-provider-order-cdl-staging-neworder
Logs:   https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fcompassdigital-provider-order-cdl-staging-neworder

LogInsights: https://console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:logs-insights$3FqueryDetail$3D$257E$2528source$257E$2528$257E$2527*2Faws*2Flambda*2Fcompassdigital-order-staging-search_index$257E$2527*2Faws*2Flambda*2Fcompassdigital-order-staging-root$257E$2527*2Faws*2Flambda*2Fcompassdigital-order-staging-order$257E$2527*2Faws*2Flambda*2Fcompassdigital-order-staging-brandorders$257E$2527*2Faws*2Flambda*2Fcompassdigital-order-staging-customerorders$257E$2527*2Faws*2Flambda*2Fcompassdigital-provider-order-cdl-staging-dexai$257E$2527*2Faws*2Flambda*2Fcompassdigital-provider-order-cdl-staging-catchall$257E$2527*2Faws*2Flambda*2Fcompassdigital-provider-order-cdl-staging-neworder$2529$2529
```