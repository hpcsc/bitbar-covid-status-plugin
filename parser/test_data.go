package parser

func GetCasesByLGAHtml() string {
	return `
<table class="table table-responsive table-bordered">
	<thead>
		<tr>
			<th scope="col">
			<p>LGA</p>
			</th>
			<th scope="col">
			<p>Confirmed cases (ever)</p>
			</th>
			<th scope="col">
			<p>Active cases (current)</p>
			</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td data-title="
			LGA
			">
			<p>WYNDHAM</p>
			</td>
			<td data-title="
			Confirmed cases (ever)
			">
			<p>1764</p>
			</td>
			<td data-title="
			Active cases (current)
			">
			<p>917</p>
			</td>
		</tr>
		<tr>
			<td data-title="
			LGA
			">
			<p>BRIMBANK</p>
			</td>
			<td data-title="
			Confirmed cases (ever)
			">
			<p>1655</p>
			</td>
			<td data-title="
			Active cases (current)
			">
			<p>876</p>
			</td>
		</tr>
		<tr>
			<td data-title="
			LGA
			">
			<p>BULOKE</p>
			</td>
			<td data-title="
			Confirmed cases (ever)
			">
			<p>0</p>
			</td>
			<td data-title="
			Active cases (current)
			">
			<p>0</p>
			</td>
		</tr>
		<tr>
			<td data-title="
			LGA
			">
			<p>HINDMARSH</p>
			</td>
			<td data-title="
			Confirmed cases (ever)
			">
			<p>0</p>
			</td>
			<td data-title="
			Active cases (current)
			">
			<p>0</p>
			</td>
		</tr>
		<tr>
			<td data-title="
			LGA
			">
			<p>TOWONG</p>
			</td>
			<td data-title="
			Confirmed cases (ever)
			">
			<p>0</p>
			</td>
			<td data-title="
			Active cases (current)
			">
			<p>0</p>
			</td>
		</tr>
		<tr>
			<td data-title="
			LGA
			">
			<p>UNKNOWN</p>
			</td>
			<td data-title="
			Confirmed cases (ever)
			">
			<p>255</p>
			</td>
			<td data-title="
			Active cases (current)
			">
			<p>202</p>
			</td>
		</tr>
		<tr>
			<td data-title="
			LGA
			">
			<p><strong>TOTAL</strong></p>
			</td>
			<td data-title="
			Confirmed cases (ever)
			">
			<p><strong>16234</strong></p>
			</td>
			<td data-title="
			Active cases (current)
			">
			<p><strong>7877</strong></p>
			</td>
		</tr>
	</tbody>
</table>`
}

func GetLatestMediaReleaseHtml() string {
	return `
<div class="page-content field field--name-field-general-paragraphs field--type-entity-reference-revisions field--label-hidden field--items">
              <div class="field--item">  <div class="paragraph paragraph--type--dhhs-rich-text paragraph--view-mode--default">
            <div class="layout layout--onecol">
    <div class="layout__region layout__region--content">
      
            <div class="field field--name-field-dhhs-rich-text-text field--type-text-long field--label-hidden field--item"><h2 id="latest-news">Latest news</h2>

<ul>
	<li><a href="/coronavirus-update-victoria-14-august-2020">Department of Health and Human Services media release - Coronavirus update for Victoria - 14&nbsp;August 2020</a></li>
	<li><a href="https://app.powerbi.com/view?r=eyJrIjoiODBmMmE3NWQtZWNlNC00OWRkLTk1NjYtMjM2YTY1MjI2NzdjIiwidCI6ImMwZTA2MDFmLTBmYWMtNDQ5Yy05Yzg4LWExMDRjNGViOWYyOCJ9">Up-to-date epidemiological data<i class="fal fa-external-link"></i></a></li>
</ul>

<h2 id="previous-media-releases">Previous media releases</h2>
</div>
      
    </div>
  </div>

      </div>
</div>
              
              
          </div>`
}

func GetLatestVictoriaNumbersHtml() string {
	return `
<div class="covid-latest-numbers">
	  	      <div class="row">
        <div class="col-sm-12">
          <h2 id="latest-victorian-numbers">Latest Victorian numbers</h2>
        </div>
      </div>
      
      <div class="row">
        <div class="col-sm-12">
          <p class="updated"><span>Updated: </span>14 August 2020   04:00pm</p>
        </div>
      </div>
      

      <div class="row covid-numbers-box equal-sizes">
        
      <div class="field field--name-field-cln-box field--type-entity-reference-revisions field--label-hidden field--items">
              <div class="field--item">

		  						  <div class="col-sm-3 col-xs-6 equal-size" style="height: 130px;">
			  	<div class="covid-number-box">

			    	<span class="numbers">372</span>

			    	<p>new cases (last 24 hours)</p>
				</div>
			  </div>
  					</div>
              <div class="field--item">

		  						  <div class="col-sm-3 col-xs-6 equal-size" style="height: 130px;">
			  	<div class="covid-number-box">

			    	<span class="numbers">7877</span>

			    	<p>active cases</p>
				</div>
			  </div>
  					</div>
              <div class="field--item">

		  						  <div class="col-sm-3 col-xs-6 equal-size" style="height: 130px;">
			  	<div class="covid-number-box">

			    	<span class="numbers">1,914,474</span>

			    	<p>total tested</p>
				</div>
			  </div>
  					</div>
              <div class="field--item">

		  						  <div class="col-sm-3 col-xs-6 equal-size" style="height: 130px;">
			  	<div class="covid-number-box">

			    	<span class="numbers">289</span>

			    	<p>lives lost</p>
				</div>
			  </div>
  					</div>
          </div>
  
      </div>

      <div class="row">
        <div class="col-sm-12 field-cln-link">
          <a href="https://app.powerbi.com/view?r=eyJrIjoiODBmMmE3NWQtZWNlNC00OWRkLTk1NjYtMjM2YTY1MjI2NzdjIiwidCI6ImMwZTA2MDFmLTBmYWMtNDQ5Yy05Yzg4LWExMDRjNGViOWYyOCJ9"><span>View full report</span> <i class="far fa-chevron-circle-right"></i><i class="fal fa-external-link"></i></a>
        </div>
      </div>

      </div>`
}

