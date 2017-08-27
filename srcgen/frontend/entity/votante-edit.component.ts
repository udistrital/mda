import { Component, OnInit } from '@angular/core';
import { Votante } from '../../models/votante';
import { VotanteService } from '../../services/votante.service';
import { Location } from '@angular/common';
import { ActivatedRoute, Params } from '@angular/router';


import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'app-votante-edit',
  templateUrl: './votante-edit.component.html',
  styleUrls: ['./votante-edit.component.css']
})
export class VotanteEditComponent implements OnInit {

  votante: Votante = new Votante();
  display = false;
  id: string;
  test = new Date('2016-01-05T09:05:05.035Z');

  constructor(private route: ActivatedRoute, private location: Location, private votanteService: VotanteService) {

  }

  actualizar(votante: Votante): void {
    this.votanteService.update(votante).then(() => this.display = true);
  }

  ngOnInit() {
    this.route.params.switchMap((params: Params) => this.votanteService.getVotante(params['id']))
      .subscribe(votante => this.votante = votante);
  }

  regresar(): void {
    this.location.back();
  }

  cerrarDialogo(): void {
    this.display = false;
    this.location.back();
  }
}