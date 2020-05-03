package vatansms

type Credential struct {
	Username string
	Password string
}

type Status struct {
	Code        int
	Description string
}

type Settings struct {
	Credential Credential
}

type Envelope struct {
	Message string
	To      string
}

type SettingsResponse struct {
	Response struct {
		Settings struct {
			Balance struct {
				Main  float64
				Limit float64
			}
			Senders []struct {
				Value   string
				Default bool
			}
			Keywords         []string
			OperatorSettings struct {
				Account   string
				MSISDN    string
				VariantId int
				ServiceId int
				UnitPrice int
			}
		}
		Status Status
	}
}

type OneToN struct {
	Credential Credential
	DataCoding string
	Header     struct {
		From                  string
		ScheduledDeliveryTime string
		ValidityPeriod        int
	}
	Message string
	To      []string
}

type NToNSms struct {
	Credential Credential
	DataCoding string
	Header     struct {
		From                  string
		ScheduledDeliveryTime string
		ValidityPeriod        int
	}
	Envelopes []Envelope
}

type SmsResponse struct {
	Response struct {
		MessageId int
		Status    Status
	}
}

type ReportStats struct {
	Credential Credential
	Range      struct {
		Begin string
		End   string
	}
}

type ReportStatsResponse struct {
	Response struct {
		Report struct {
			List []struct {
				Year        int
				Month       int
				Delivered   int
				Undelivered int
				Count       int
			}
		}
		Status Status
	}
}

type ReportDetails struct {
	Credential Credential
	MessageId  int
	MSISDN     string
}

type ReportDetailsResponse struct {
	Response struct {
		ReportDetail struct {
			List []struct {
				Id          int
				Network     int
				MSISDN      string
				Cost        float64
				Submitted   string
				LastUpdated string
				State       string
				Sequence    int
				ErrorCode   int
				Payload     string
				Xser        string
			}
		}
		Status Status
	}
}
