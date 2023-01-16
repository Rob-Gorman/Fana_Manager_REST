import moment from "moment"

export const flags = [
  {
    _id: '123',
    key: "flag_1",
    displayName: "Flag 1",
    status: false,
    audiences: [{ name: "Beta testers" }, { name: "California Students" }],
    combine: "ALL",
    createdAt: moment().format("MMM Do YY"),
    updatedAt: moment().format("MMM Do YY")
  },
  {
    _id: '124',
    key: "flag_2",
    displayName: "Flag 2",
    status: true,
    audiences: [{ name: "Beta testers" }],
    combine: "ANY",
    createdAt: moment().format("MMM Do YY"),
    updatedAt: moment().format("MMM Do YY")
  }
]

export const audiences = [
  {
    combine: "ANY",
    conditions: [
      {
        attribute: 'beta',
        operator: 'EQ',
        value: 'true'
      }],
    createdAt: moment().format("MMM Do YY"),
    updatedAt: moment().format("MMM Do YY"),
    _id: 'abc',
    key: "beta_testers",
    displayName: "Beta Testers"
  },
  {
    combine: "ALL",
    conditions: [
      {
        attribute: 'state',
        operator: 'EQ',
        value: 'california'
      }
    ],
    createdAt: moment().format("MMM Do YY"),
    updatedAt: moment().format("MMM Do YY"),
    _id: 'abd',
    key: "california_students",
    displayName: "California Students"
  },
]

export const attributes = [
  {
    _id: "1",
    name: "beta",
    attrType: "BOOL",
    type: "boolean"
  },
  {
    _id: "2",
    name: "state",
    attrType: "STR",
    type: "string"
  },
  {
    _id: "3",
    name: "userId",
    attrType: "STR",
    type: "string"
  }
]