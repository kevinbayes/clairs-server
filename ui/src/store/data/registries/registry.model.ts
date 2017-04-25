export namespace Registries {

  export class NewRegistry {
    name: string;
    description: string;
    uri: string;

    credentials: Credentials;
  }

  export class Credentials {
    username: string;
    password: string;
  }
}
