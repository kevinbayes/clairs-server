import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { RegistryContainersComponent } from './registry.containers.component';

describe('ContainersComponent', () => {
  let component: RegistryContainersComponent;
  let fixture: ComponentFixture<RegistryContainersComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ RegistryContainersComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(RegistryContainersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
