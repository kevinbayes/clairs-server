import { TestBed, inject } from '@angular/core/testing';

import { RegistriesService } from './registries.service';

describe('RegistriesService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [RegistriesService]
    });
  });

  it('should ...', inject([RegistriesService], (service: RegistriesService) => {
    expect(service).toBeTruthy();
  }));
});
