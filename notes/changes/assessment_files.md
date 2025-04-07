
# Proposed Change: Store Assessment Instructions as Markdown in Database instead of using File field

## Status

Pending development

## Summary

Store assessment instructions as markdown in instructions column of assessment row.

## Justification

- Added flexibility and simplicity: Easier to move assessments around. Filesystem will not need to be part of assessment CRUD operations. Faster read and write

## Negative Impact

- A few hours of time (several days of work)

## Required actions

- [ ] Create migration to write all existing assessment file content to instructions column of corresponding row
- [ ] Remove file field from assessment struct
- [ ] Delete assessment files
- [ ] Teacher view (echo server):
  - [ ] Remove view file content UI
  - [ ] Remove file name UI
  - [ ] Assessments UI: button to view instructions as html in iframe or new tab instead of regular text
  - [ ] Assessments UI: button to edit instructions as markdown
  - [ ] Assessments service method to render instructions markdown content to html using goldmark
  - [ ] Assessments service method to edit instructions as markdown in markdown editor
- [ ] Student view (static site generator): Assessments UI to view instructions as html
