import {AfterViewInit, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {
  IDatatablePaginationEvent,
  IDatatableSelectionEvent, IDatatableSortEvent, MdDataTableComponent,
  MdDataTablePaginationComponent
} from "ng2-md-datatable";
import {Observable} from "rxjs/Observable";
import 'rxjs/add/observable/from';
import {Subject} from "rxjs/Subject";
import {RegistriesService} from "../../services/registries.service";
import {MdDialog} from "@angular/material";
import {NewRegistryModalComponent} from "./new/new.registry.modal.component";

@Component({
  selector: 'app-registries',
  templateUrl: './registries.component.html',
  styleUrls: ['./registries.component.less']
})
export class RegistriesComponent implements OnInit, AfterViewInit, OnDestroy {

  @ViewChild(MdDataTableComponent) datatable: MdDataTableComponent;
  @ViewChild(MdDataTablePaginationComponent) pager: MdDataTablePaginationComponent;

  public pagination: any = {
    Pages: 0,
    Page: 0,
    Size: 10
  };
  public registries: any[];
  private unmount$: Subject<void> = new Subject<void>();

  constructor(
    private registriesService: RegistriesService,
    public dialog: MdDialog) { }

  ngOnInit() {

    this.registriesService.all().subscribe((res) => {

      console.log(res);
      this.pagination = res.Meta;
      this.registries = res.Entities;
    }, (err) => {

      console.error(err);
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
    let dialogRef = this.dialog.open(NewRegistryModalComponent);
      dialogRef.afterClosed().subscribe(result => {
    });
  }

  ngOnDestroy() {
    this.unmount$.next();
    this.unmount$.complete();
  }

}
