import {Inject, Injectable} from '@angular/core';
import {Http, Response} from "@angular/http";
import {Observable} from "rxjs/Observable";
import {Registries} from "../../store/data/registries/registry.model";
import NewRegistry = Registries.NewRegistry;

@Injectable()
export class RegistriesService {

  baseUrl: string;

  constructor(@Inject(Http) private http: Http) {

    this.baseUrl = "/api/registries";
  }

  public save(registry: NewRegistry): Observable<any> {

    return this.http
      .post(`${this.baseUrl}`, registry);
  }

  public update(registry: any): Observable<any> {

    return this.http
      .put(`${this.baseUrl}/${registry.Id}`, registry);
  }

  public listContainers(registry: any, page: number = 0, size:number = 10): Observable<any> {

    return this.http
      .get(`${this.baseUrl}/${registry}/containers`)
      .map(r => r.json());
  }

  public all(page: number = 0, size:number = 10, _container_count: boolean = true): Observable<any> {

    var url = (_container_count ? `/api/summary/registries?p=${page}&s=${size}` : `${this.baseUrl}?p=${page}&s=${size}`)

    return this.http
      .get(url)
      .map(r => r.json());
  }
}
