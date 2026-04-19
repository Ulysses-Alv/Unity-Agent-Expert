# Verified Solutions onboarding process

Before the Verified Solutions team can start technical verification, you must submit the following information through [this form](https://airtable.com/shrrTKYkeNzeAEitJ):

* A link to the solution.
* Size of the solution.
* Supported Unity versions.
* Supported platforms.
* Any necessary licensing files.
* Login information to any back-end services or administration consoles.
* A link to the documentation.

You must also complete detailed internal QA procedures.

## Custom EULA

You can distribute a Verified Solution through the Unity **Asset Store** with either the standard Asset Store end user license agreement (EULA) or a custom EULA. The standard Unity Asset Store EULA is appropriate for asset type packages, including 3D and **2D objects**, game templates, and audio packs. Tool- or SDK-type packages may require specific additional information in the EULA.

To include a custom EULA, add a sentence to the description of the solution stating the following: `This asset is governed by the <provider's name> EULA <link to custom EULA>.` You must email the Verified Solutions team at `vs-support@unity3d.com` regarding the custom EULA.

## Technical verification process

The Verified Solutions team evaluates solutions in the following stages:

* [Exploration](#exploration): Understands the solution.
* [Verification](#verification): Tests each feature of the solution.
* [Report](#report): Provides a report of their analysis.

### Exploration

The Verified Solutions team dedicates the beginning of the verification process to understanding the solution. The team evaluates the solution to determine its size and complexity. The team uses documentation and other available learning resources (such as demos and tutorial videos) to understand the solution and determine the core and most commonly used features. The team uses this information to set up a new Unity project for the solution.

### Verification

During this step, the Verified Solutions team creates automated tests to cover all core APIs and help with any other areas that might require manual work. The team uses solution scope, defined during the exploration phase, to determine whether more obscure APIs and features can be reasonably tested in as much detail as the main features.

To make sure customers can trust the solution to work on any of the supported platforms, the solution is tested on each platform, beginning with one Unity version. If the team finds an issue that prevents use of the solution, they stop verification and inform you of the issue. Otherwise, the solution is then tested on all Unity LTS streams and the latest Tech stream. The team verifies documentation during this step to make sure that customers can find all the needed information regardless of their level of experience with Unity services.

During the final step of the verification process, the Verified Solutions team verifies the solution using the [Asset Store Tools](https://assetstore.unity.com/packages/package/5368745) Validator. Although the Verified Solutions program has a custom end user license agreement (EULA), you must verify that the solution you’re submitting follows all applicable guidelines for Asset Store Tools. The Verified Solutions team documents any issues found during the verification process and details them during the [report](#report) phase.

### Report

After the Verified Solutions team tests the solution, they submit a verification report that details insights found during the process.

The first page of the report details the core information about the solution, and the summary section of the report includes:

* An overview of past verifications.
* Exact Unity versions used.
* Test device information.
* Covered platforms.
* Test statistics.

This section provides an overview of the outcome of the report and highlights important issues alongside other notes and suggestions from Unity that can improve the user experience.

![The first section of the verification report includes information about submission dates, data, and testing summaries for the verified solution.](../uploads/Main/VerifiedSolutionsReportExample1.png)

The first section of the verification report includes information about submission dates, data, and testing summaries for the verified solution.

The next section of the report contains the detailed information on each test case:

![The second section of the verification report includes an itemized list of the performed tests, whether the solution passed or failed these tests, and comments and bug severity ratings from the Unity testing team.](../uploads/Main/VerifiedSolutionsReportExample2.png)

The second section of the verification report includes an itemized list of the performed tests, whether the solution passed or failed these tests, and comments and bug severity ratings from the Unity testing team.

The report categorizes issues found during the verification phase by severity using the following categories:

* **Minor**: Non-core or niche features don’t function as expected. For example, warnings aren’t user-friendly, non-breaking errors or exceptions display, or the feature has cosmetic issues that don’t prevent the usability of the feature.
* **Major**: Core solution features don’t function as expected. For example, the Unity Editor or Unity Player crashes, which results in a loss of progress or data, or errors are commonly and continuously encountered.
* **Critical**: The solution is unusable due to issues. For example, the solution doesn’t work with primary target platforms, or the solution includes issues that risk causing problems to the operating system and Unity products.

To receive approval as a Verified Solution, a product can’t contain any major or critical issues.

Other Verified Solutions team members review the completed verification report to ensure all issues are categorized correctly and that the report is accurate and complete.

If the Verified Solutions team found major or critical issues in the solution during the verification process, you must fix these issues before the Verified Solutions team can reverify the solution. When the team completes verification, they send the solution to the Verified Solutions publisher, and the product is officially recognized as a Verified Solution.

## Additional resources

* [Process to become a Verified Solution](verified-solutions-process.md)
* [Verified Solutions program](https://assetstore.unity.com/verified-solutions)