## www.vatansms.net Golang Paketi
### Hesap bilgileri ve kredi sorgulamak için;
    settings := vatansms.Settings{}
    settings.Credential.Username = "vatan"
    settings.Credential.Password = "123456"
    settingsResponse,err := settings.GetSettings()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(settingsResponse)
   
### Bir mesajı N numaraya göndermek için;
    oneToNSms := vatansms.OneToN{}
    oneToNSms.Credential.Username = "vatan"
    oneToNSms.Credential.Password = "123456"
    oneToNSms.Message = "merhaba"
    oneToNSms.Header.From = "VATANSMS"

    //Bu parametre sadece ileri tarihli mesajlarda gönderilmeli. Format : 2006-01-02T15:04:05-0700
    //oneToNSms.Header.ScheduledDeliveryTime = fiveMinutesLater

    //Türkçe karakterli SMS gönderimi için DataCoding değerini TurkishSingleShift girmelisiniz.
    oneToNSms.DataCoding = "DEFAULT"
    oneToNSms.To = []string{vatansms.PhoneVerify("5356992106")}
    oneToNResponse,err := oneToNSms.SendOneToN()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(oneToNResponse)
    
### N mesajı N numaraya göndermek için;
    nToNSms := vatansms.NToNSms{}
    nToNSms.Credential.Username = "vatan"
    nToNSms.Credential.Password = "123456"
    nToNSms.Header.From = "VATANSMS"

    //Bu parametre sadece ileri tarihli mesajlarda gönderilmeli. Format : 2006-01-02T15:04:05-0700
    //nToNSms.Header.ScheduledDeliveryTime = fiveMinutesLater

    //Türkçe karakterli SMS gönderimi için DataCoding değerini TurkishSingleShift girmelisiniz.
    nToNSms.DataCoding = "DEFAULT"
    nToNSms.Envelopes = append(nToNSms.Envelopes, vatansms.Envelope{
        Message: "Merhaba 1",
        To:      vatansms.PhoneVerify("5356992106"),
    })
    nToNSms.Envelopes = append(nToNSms.Envelopes, vatansms.Envelope{
        Message: "Merhaba 2",
        To:      vatansms.PhoneVerify("5996662121"),
    })
    nToNResponse, err := nToNSms.SendNToNSms()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(nToNResponse)      


### Tarih aralığı ile mesajların sonuç toplamları için;
    reportStats := vatansms.ReportStats{}
	reportStats.Credential.Username = "vatan"
	reportStats.Credential.Password = "123456"
	reportStats.Range.Begin = time.Now().AddDate(0,0,-1).Format("2006-01-02T15:04:05-0700")
	reportStats.Range.End = time.Now().Format("2006-01-02T15:04:05-0700")
	reportStatsResponse,err := reportStats.GetReportStats()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(reportStatsResponse)   


### Mesajın detaylı olarak numaralara ulaştı/ulaşmadı durumunu sorgulamak için;
    reportDetails := vatansms.ReportDetails{}
	reportDetails.Credential.Username = "vatan"
	reportDetails.Credential.Password = "123456"
	reportDetails.MessageId = 1142477679

	// Eğer mesajın içerisindeki sadece bir numaraya ait rapor isteniyorsa, bu durumda numara aşağıdaki parametre ile bildirilmeli
	//reportDetails.MSISDN = vatansms.PhoneVerify("5356992106")
	reportDetailsResponse, err := reportDetails.GetReportDetails()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(reportDetailsResponse)
	
	