import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NewRegistryModalComponent } from './new.registry.modal.component';

describe('NewComponent', () => {
  let component: NewRegistryModalComponent;
  let fixture: ComponentFixture<NewRegistryModalComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NewRegistryModalComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NewRegistryModalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
