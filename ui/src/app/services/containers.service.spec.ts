import { TestBed, inject } from '@angular/core/testing';

import { ContainersService } from './containers.service';

describe('ContainersService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ContainersService]
    });
  });

  it('should ...', inject([ContainersService], (service: ContainersService) => {
    expect(service).toBeTruthy();
  }));
});
