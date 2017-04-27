import {AfterViewInit, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {ContainersService} from "../../../services/containers.service";
import {
  IDatatablePaginationEvent,
  IDatatableSelectionEvent, IDatatableSortEvent, MdDataTableComponent,
  MdDataTablePaginationComponent
} from "ng2-md-datatable";
import {Observable} from "rxjs/Observable";
import 'rxjs/add/observable/from';
import {Subject} from "rxjs/Subject";
import {RegistryNewContainerModalComponent} from "./new/registry.new.container.modal.component";
import {MdDialog} from "@angular/material";
import {RegistriesService} from "../../../services/registries.service";
import {ActivatedRoute, Route} from "@angular/router";

@Component({
  selector: 'app-registry-containers',
  templateUrl: './registry.containers.component.html',
  styleUrls: ['./registry.containers.component.less']
})
export class RegistryContainersComponent implements OnInit, AfterViewInit, OnDestroy {

  @ViewChild(MdDataTableComponent) datatable: MdDataTableComponent;
  @ViewChild(MdDataTablePaginationComponent) pager: MdDataTablePaginationComponent;

  private id: number;

  public pagination: any = {
    Pages: 0,
    Page: 0,
    Size: 10
  };
  public containers: any[];
  private unmount$: Subject<void> = new Subject<void>();

  constructor(private containersService: ContainersService,
              private registriesService: RegistriesService,
              private route: ActivatedRoute,
              public dialog: MdDialog) { }

  ngOnInit() {

    this.route.params.subscribe(params => {

      this.id = +params['id']; // (+) converts string 'id' to a number

      this.registriesService.listContainers(this.id).subscribe((res) => {
        this.pagination = res.Meta;
        this.containers = res.Entities;
      }, (err) => {

        console.error(err);
      });
    });
  }

  ngAfterViewInit() {
    if (this.datatable) {
      Observable.from(this.datatable.selectionChange)
        .takeUntil(this.unmount$)
        .subscribe((e: IDatatableSelectionEvent) => console.log("Data selected"));

      Observable.from(this.datatable.sortChange)
        .takeUntil(this.unmount$)
        .subscribe((e: IDatatableSortEvent) =>
          console.log("Sort order"));

      Observable.from(this.pager.paginationChange)
        .takeUntil(this.unmount$)
        .subscribe((e: IDatatablePaginationEvent) =>
          console.log("Pager changed"));
    }
  }

  openDialog() {
    let dialogRef = this.dialog.open(RegistryNewContainerModalComponent);
    dialogRef.afterClosed().subscribe(result => {
    });
  }

  ngOnDestroy() {
    this.unmount$.next();
    this.unmount$.complete();
  }

}
