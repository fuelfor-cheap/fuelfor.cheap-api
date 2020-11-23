class InvalidSchema {
  error: boolean;

  message: string;

  schema: string;

  constructor(schema) {
    this.error = true;
    this.message = 'malformed payload details, request body should follow the follwing schema';
    this.schema = schema;

    return this;
  }
}

export function invalidSchema(schema: any): InvalidSchema {
  return new InvalidSchema(JSON.stringify(schema));
}