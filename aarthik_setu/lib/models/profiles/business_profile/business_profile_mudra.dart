import '../../common/bank_details.dart';
import 'components/manual_itr_business.dart';
import 'components/key_person_details_mudra.dart';

class BusinessProfile{
  Map<(int, int), String> itr;
  Map<(int, int), ManualItrBusiness> manualItr;
  List<BankDetails> bankDetails;
  KeyPersonDetails keyPersonDetails;

  BusinessProfile({
    required this.itr,
    required this.manualItr,
    required this.bankDetails,
    required this.keyPersonDetails,
  });

  factory BusinessProfile.fromJson(Map<String, dynamic> json) {
    return BusinessProfile(
      itr: json['itr'],
      manualItr: json['manualItr'],
      bankDetails: json['bankDetails'],
      keyPersonDetails: json['keyPersonDetails'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'itr': itr,
      'manualItr': manualItr,
      'bankDetails': bankDetails,
      'keyPersonDetails': keyPersonDetails,
    };
  }
}