---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_refresh_schedule"
description: |-
  Manages a Resource QuickSight Refresh Schedule.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_quicksight_refresh_schedule

Resource for managing a QuickSight Refresh Schedule.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { QuicksightRefreshSchedule } from "./.gen/providers/aws/quicksight-refresh-schedule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new QuicksightRefreshSchedule(this, "example", {
      dataSetId: "dataset-id",
      schedule: [
        {
          refreshType: "FULL_REFRESH",
          scheduleFrequency: [
            {
              interval: "HOURLY",
            },
          ],
        },
      ],
      scheduleId: "schedule-id",
    });
  }
}

```

### With Weekly Refresh

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { QuicksightRefreshSchedule } from "./.gen/providers/aws/quicksight-refresh-schedule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new QuicksightRefreshSchedule(this, "example", {
      dataSetId: "dataset-id",
      schedule: [
        {
          refreshType: "INCREMENTAL_REFRESH",
          scheduleFrequency: [
            {
              interval: "WEEKLY",
              refreshOnDay: [
                {
                  dayOfWeek: "MONDAY",
                },
              ],
              timeOfTheDay: "01:00",
              timezone: "Europe/London",
            },
          ],
        },
      ],
      scheduleId: "schedule-id",
    });
  }
}

```

### With Monthly Refresh

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { QuicksightRefreshSchedule } from "./.gen/providers/aws/quicksight-refresh-schedule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new QuicksightRefreshSchedule(this, "example", {
      dataSetId: "dataset-id",
      schedule: [
        {
          refreshType: "INCREMENTAL_REFRESH",
          scheduleFrequency: [
            {
              interval: "MONTHLY",
              refreshOnDay: [
                {
                  dayOfMonth: "1",
                },
              ],
              timeOfTheDay: "01:00",
              timezone: "Europe/London",
            },
          ],
        },
      ],
      scheduleId: "schedule-id",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `dataSetId` - (Required, Forces new resource) The ID of the dataset.
* `scheduleId` - (Required, Forces new resource) The ID of the refresh schedule.
* `schedule` - (Required) The [refresh schedule](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_RefreshSchedule.html). See [schedule](#schedule)

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `awsAccountId` - (Optional, Forces new resource) AWS account ID.

### schedule

* `refreshType` - (Required) The type of refresh that the dataset undergoes. Valid values are `INCREMENTAL_REFRESH` and `FULL_REFRESH`.
* `startAfterDateTime` (Optional) Time after which the refresh schedule can be started, expressed in `YYYY-MM-DDTHH:MM:SS` format.
* `scheduleFrequency` - (Optional) The configuration of the [schedule frequency](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_RefreshFrequency.html). See [schedule_frequency](#schedule_frequency).

### schedule_frequency

* `interval` - (Required) The interval between scheduled refreshes. Valid values are `MINUTE15`, `MINUTE30`, `HOURLY`, `DAILY`, `WEEKLY` and `MONTHLY`.
* `timeOfTheDay` - (Optional) The time of day that you want the dataset to refresh. This value is expressed in `HH:MM` format. This field is not required for schedules that refresh hourly.
* `timezone` - (Optional) The timezone that you want the refresh schedule to use.
* `refreshOnDay` - (Optional) The [refresh on entity](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ScheduleRefreshOnEntity.html) configuration for weekly or monthly schedules. See [refresh_on_day](#refresh_on_day).

### refresh_on_day

* `dayOfMonth` - (Optional) The day of the month that you want to schedule refresh on.
* `dayOfWeek` - (Optional) The day of the week that you want to schedule a refresh on. Valid values are `SUNDAY`, `MONDAY`, `TUESDAY`, `WEDNESDAY`, `THURSDAY`, `FRIDAY` and `SATURDAY`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Amazon Resource Name (ARN) of the refresh schedule.
* `id` - A comma-delimited string joining AWS account ID, data set ID & refresh schedule ID.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import a QuickSight Refresh Schedule using the AWS account ID, data set ID and schedule ID separated by commas (`,`). For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { QuicksightRefreshSchedule } from "./.gen/providers/aws/quicksight-refresh-schedule";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    QuicksightRefreshSchedule.generateConfigForImport(
      this,
      "example",
      "123456789012,dataset-id,schedule-id"
    );
  }
}

```

Using `terraform import`, import a QuickSight Refresh Schedule using the AWS account ID, data set ID and schedule ID separated by commas (`,`). For example:

```console
% terraform import aws_quicksight_refresh_schedule.example 123456789012,dataset-id,schedule-id
```

<!-- cache-key: cdktf-0.20.8 input-df51f205a14415df4f456a34719514085a31ad470b6e47d269510c41210a7783 -->