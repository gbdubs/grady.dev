## To Run Locally

cd hugodata; hugo server --watch --disableFastRender; 

visit http://localhost:1313/project

## To deploy

just commit to main, the `build-for-ghpages` action will do some post-processing,
including running hugo, and commit it to the `gh-pages` branch, which
github will then pickup and push.