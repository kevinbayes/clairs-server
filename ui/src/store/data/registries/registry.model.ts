export namespace Registries {

  export class BaseRegistry {
    name: string;
    description: string;
    uri: string;

    credentials: Credentials;
  }

  export class Credentials {
    username: string;
    password: string;
  }

  export class Registry extends BaseRegistry {
    id: number;

  }
}
