export const operators = {
  EQ: '=',
  IN: 'is in',
  STR_CONTAINS: 'contains',
  STR_STARTS_WITH: 'starts with',
  STR_ENDS_WITH: 'ends with',
  GT: '>',
  LT: '<',
  LT_EQ: '<=',
  GT_EQ: '>='
}

export const combinationOptions = [
  { value: '', text: 'Select a type' },
  { value: 'ANY', text: 'Any' },
  { value: 'ALL', text: 'All' },
];

export const operatorOptions = [
  { value: '', text: 'Select an operator' },
  { value: 'EQ', text: 'is equal to' },
  { value: 'IN', text: 'is in' },
  { value: 'NOT_IN', text: 'is not in' },
  { value: 'STR_CONTAINS', text: 'contains' },
  { value: 'STR_STARTS_WITH', text: 'starts with' },
  { value: 'STR_ENDS_WITH', text: 'ends with' },
  { value: 'GT', text: '>' },
  { value: 'LT', text: '<' },
  { value: 'GT_EQ', text: '>=' },
  { value: 'LT_EQ', text: '<=' },
];

export const operatorOptionsByType = {
  STR: [
    { value: 'EQ', text: 'is equal to' },
    { value: 'IN', text: 'is in' },
    { value: 'STR_CONTAINS', text: 'contains' },
    { value: 'STR_STARTS_WITH', text: 'starts with' },
    { value: 'STR_ENDS_WITH', text: 'ends with' },
  ],
  BOOL: [
    { value: 'EQ', text: 'is equal to' },
  ],
  NUM: [
    { value: 'EQ', text: 'is equal to' },
    { value: 'GT', text: '>' },
    { value: 'LT', text: '<' },
    { value: 'GT_EQ', text: '>=' },
    { value: 'LT_EQ', text: '<=' },
  ]
}


export  const attrTypeOptions = [
  {value: "", text:"Select a type"}, 
  {value: "BOOL", text:"Boolean"}, 
  {value: "STR", text: "String"}, 
  {value: "NUM", text: "Number"}, 
  {value: "DATETIME", text: "Date/Timestamp"}
]

export const attrTypeMapper = {
  BOOL: 'Boolean',
  STR: 'String',
  NUM: 'Number',
};