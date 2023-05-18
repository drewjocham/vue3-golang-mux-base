export interface TestResponse {
    firstName: string
    lastName: string
}

export interface TestResponseRO {
    data: TestResponse[]
    metadata: string
}
