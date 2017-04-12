/**
 * Created by kevinbayes on 13/04/17.
 */
import {
  MdButtonModule,
  MdCardModule,
  MdCheckboxModule,
  MdGridListModule,
  MdIconModule,
  MdInputModule,
  MdListModule,
  MdMenuModule,
  MdProgressSpinnerModule,
  MdRippleModule,
  MdSidenavModule,
  MdSnackBarModule,
  MdTabsModule,
  MdToolbarModule
} from '@angular/material';
import { NgModule } from "@angular/core";

@NgModule({
  imports:      [
    MdButtonModule,
    MdCheckboxModule,
    MdToolbarModule,
    MdCardModule,
    MdMenuModule,
    MdIconModule,
    MdInputModule,
    MdGridListModule,
    MdListModule,
    MdProgressSpinnerModule,
    MdSnackBarModule,
    MdRippleModule,
    MdSidenavModule,
    MdTabsModule
  ],
  exports:      [
    MdButtonModule,
    MdCheckboxModule,
    MdToolbarModule,
    MdCardModule,
    MdMenuModule,
    MdIconModule,
    MdInputModule,
    MdGridListModule,
    MdListModule,
    MdProgressSpinnerModule,
    MdSnackBarModule,
    MdRippleModule,
    MdSidenavModule,
    MdTabsModule
  ]
})
export class MaterialModule { }
