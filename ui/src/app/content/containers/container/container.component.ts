import {AfterViewInit, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {ContainersService} from "../../../services/containers.service";
import {Subject} from "rxjs/Subject";
import {Observable} from "rxjs/Observable";
import {
  IDatatablePaginationEvent, IDatatableSelectionEvent, IDatatableSortEvent,
  MdDataTableComponent, MdDataTablePaginationComponent
} from "ng2-md-datatable";
import {ActivatedRoute} from "@angular/router";

@Component({
  selector: 'app-container',
  templateUrl: './container.component.html',
  styleUrls: ['./container.component.less']
})
export class ContainerComponent implements OnInit, AfterViewInit, OnDestroy {

  @ViewChild(MdDataTableComponent) datatable: MdDataTableComponent;
  @ViewChild(MdDataTablePaginationComponent) pager: MdDataTablePaginationComponent;

  private id: number;
  public pagination: any = {
    Pages: 0,
    Page: 0,
    Size: 10
  };
  public container: any = {Tags:[]};
  private unmount$: Subject<void> = new Subject<void>();

  constructor(private route: ActivatedRoute, private containersService: ContainersService) { }

  ngOnInit() {

    this.route.params.subscribe(params => {

      this.id = +params['id']; // (+) converts string 'id' to a number

      this.containersService.get(this.id).subscribe((res) => {

        console.log(res);
        this.container = res.Entity;
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

  ngOnDestroy() {
    this.unmount$.next();
    this.unmount$.complete();
  }

}
