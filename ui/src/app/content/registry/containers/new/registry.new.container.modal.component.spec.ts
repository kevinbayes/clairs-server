import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { RegistryNewContainerModalComponent } from './registry.new.container.modal.component';

describe('NewComponent', () => {
  let component: RegistryNewContainerModalComponent;
  let fixture: ComponentFixture<RegistryNewContainerModalComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ RegistryNewContainerModalComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(RegistryNewContainerModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
