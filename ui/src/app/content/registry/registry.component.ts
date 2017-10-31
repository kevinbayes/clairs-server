import { Component, OnInit } from '@angular/core';
import {RegistriesService} from "../../services/registries.service";
import {ActivatedRoute} from "@angular/router";

@Component({
  selector: 'app-registry',
  templateUrl: './registry.component.html',
  styleUrls: ['./registry.component.less']
})
export class RegistryComponent implements OnInit {

  private registry: any = {};
  private dashboard: any = {};
  private id: number;

  constructor(private registryService: RegistriesService,
              private route: ActivatedRoute) { }

  ngOnInit() {

    this.route.params.subscribe(params => {

      this.id = +params['id'];
      this.loadDashboard();
    });
  }

  loadDashboard() {

    this.registryService.dashboard(this.id)
      .subscribe((success) => {

        this.registry = success.Entity.Registry;
        this.dashboard = success.Entity;
      }, (err) => {

        console.log("Error loading the dashboard.");
        console.error(err);
      });
  }
}
